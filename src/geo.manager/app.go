package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	core "core"
	natsTopic "nats/topic"

	// Имена топиков
	natsTopicManager "nats/topic/manager"
	natsTopicWorker "nats/topic/worker"

	"geo.manager/config"
	coreBase "geo.manager/core/base"
	coreComponents "geo.manager/core/components"
	corePubs "geo.manager/core/pubs"

	stModels "geo.manager/structs/models"

	// Обработчики cron-задач
	cronCitiesUpdateTTL "geo.manager/cron/cities/update_ttl"
	cronTaskSearch "geo.manager/cron/task/search"
	cronWorkerCheckRegister "geo.manager/cron/worker/check_register"
	cronWorkerUpdateTTL "geo.manager/cron/worker/update_ttl"

	// Обработкичи http-запросов
	"geo.manager/entrypoints/api/city/add_city"
	"geo.manager/entrypoints/api/ping"

	// Обработчики событий
	eventCitiesCities "geo.manager/event/cities/cities"
	eventTaskWorkerSearch "geo.manager/event/task/worker/search"
	eventWorkerRegister "geo.manager/event/worker/register"
	eventWorkerWorkers "geo.manager/event/worker/workers"

	// Указываем свои каналы связи
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/transport/nats"

	// DB
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/encoder"
	"github.com/micro/go-micro"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

// Список сервисов сети geo, которые использует текущий сервис
const (
	GeoWorker = "geo.worker"
)

// Структура приложения
type App struct {
	core.App

	configuration *config.Configuration
	db            *gorm.DB
	webServer     *martini.ClassicMartini
	service       micro.Service
	com           *coreComponents.Components
	pub           *corePubs.Pubs
	base          *coreBase.Base
}

// Инициализация брокера для запросов City.Cities
func (a *App) initCityCitiesPubSub() {
	natsTopic.NewSubs(a.configuration.Subscrption.Workers, a.configuration.App, natsTopicManager.CityResponseStations, a.service, func() interface{} {
		return eventCitiesCities.NewCities(a.com)
	})

	a.pub.PubWorkerCityStations = natsTopic.NewPub(GeoWorker, natsTopicWorker.CityRequestStations, a.service)
}

// Инициализация брокера для запросов Task.Worker.Search
func (a *App) initTaskWorkerSearchPubSub() {
	natsTopic.NewSubs(a.configuration.Subscrption.Workers, a.configuration.App, natsTopicManager.TaskWorkerResponseSearch, a.service, func() interface{} {
		return eventTaskWorkerSearch.NewResponse(a.com)
	})

	natsTopic.NewSubs(a.configuration.Subscrption.Workers, a.configuration.App, natsTopicManager.TaskWorkerLogSearch, a.service, func() interface{} {
		return eventTaskWorkerSearch.NewLog(a.com)
	})
}

// Инициализация подписчика для получения регистрационных данных воркера
func (a *App) initSubWorkerRegister() {
	natsTopic.NewSubs(a.configuration.Subscrption.Workers, a.configuration.App, natsTopicManager.WorkerEventRegister, a.service, func() interface{} {
		return eventWorkerRegister.NewRegister(a.com)
	})
}

// Инициализация подписчика для получения данных воркеров от воркера
func (a *App) initSubWorkerWorkers() {
	natsTopic.NewSubs(a.configuration.Subscrption.Workers, a.configuration.App, natsTopicManager.WorkerEventWorkers, a.service, func() interface{} {
		return eventWorkerWorkers.NewWorkers(a.com)
	})
}

// Инициализация micro
func (a *App) initMicro() {
	// Готовим переменное окружение для micro
	os.Setenv("MICRO_REGISTRY", a.configuration.Micro.Registry)
	os.Setenv("MICRO_REGISTRY_ADDRESS", a.configuration.Micro.RegistryAddress)
	os.Setenv("MICRO_BROKER", a.configuration.Micro.Broker)
	os.Setenv("MICRO_BROKER_ADDRESS", a.configuration.Micro.BrokerAddress)
	os.Setenv("MICRO_TRANSPORT", a.configuration.Micro.Transport)
	os.Setenv("MICRO_TRANSPORT_ADDRESS", a.configuration.Micro.TransportAddress)

	// Экземпляр сервиса
	a.service = micro.NewService(
		micro.Name(a.configuration.App.Name),
		micro.Version(a.configuration.App.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	a.service.Init()
}

// Инициализация работы с БД
func (a *App) initDB() {
	var err error

	// Коннектимся к базе
	a.db, err = gorm.Open("mysql", a.configuration.Database.Username+":"+a.configuration.Database.Password+"@tcp("+a.configuration.Database.Host+":"+strconv.Itoa(a.configuration.Database.Port)+")/"+a.configuration.Database.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	if err := a.db.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	// By https://stackoverflow.com/questions/28135580/golang-mysql-error-1040-too-many-connections
	// Ограничение по количеству соединений с базой
	//a.db.DB().SetMaxOpenConns(1000)
	//a.db.DB().SetConnMaxLifetime(0)

	//a.db.LogMode(true)
	a.db.DB().SetMaxIdleConns(100)
	a.db.DB().SetMaxOpenConns(600)

	// Устанавливаем префикс таблицам
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "eb_" + defaultTableName
	}

	// Миграция схем таблиц
	a.db.AutoMigrate(
		&stModels.CitiesLogs{},
		&stModels.Cities{},
		&stModels.Workers{},
	)

	// Указываем свой логгер для GORM
	//a.db.SetLogger(log.New())
}

// Отключаем логгирование у мартини
func (a *App) classicWithoutLogging() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(martini.Recovery())
	m.Use(martini.Static("public"))
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)

	return &martini.ClassicMartini{m, r}
}

// Инициализация веб-сервера
func (a *App) initWebServer() {
	// Включаем прод режим
	martini.Env = martini.Prod

	// Инициализация сервера Мартини
	a.webServer = a.classicWithoutLogging()

	// Настройка "middleware"
	// Сервис для представления данных в нескольких форматах и взаимодействия с контентом
	a.webServer.Use(func(c martini.Context, w http.ResponseWriter, r *http.Request) {
		c.MapTo(encoder.JsonEncoder{}, (*encoder.Encoder)(nil))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	// Роутинг запросов
	a.webServer.Get("/api/ping/", ping.NewPing().Handler)

	// Маппинг конфига
	a.webServer.Map(a.configuration)

	// Внедрим в Martini базу данных
	// Сервис будет доступен для всех обработчиков как *sqlx.DB
	a.webServer.Map(a.db)
}

// Инициализация планировщика заданий
func (a *App) initCron() {
	// Аля cron на горутинах
	// by https://github.com/robfig/cron + https://stackoverflow.com/questions/16466320/is-there-a-way-to-do-repetitive-tasks-at-intervals-in-golang
	cron := cron.New()

	// Будем каждые 1сек слать пробивать с города
	if err := cron.AddFunc("* * * * * *", cronTaskSearch.NewSearch(a.com).Run); err != nil {
		log.Fatal(err)
	}

	// Будем каждые 10сек проверять актуальности зарегистрированных воркеров по их TTL
	if err := cron.AddFunc("*/10 * * * * *", cronWorkerCheckRegister.NewCheckRegister(a.com).Run); err != nil {
		log.Fatal(err)
	}

	// Будем каждые 1сек пересчитывать TTL воркеров
	if err := cron.AddFunc("* * * * * *", cronWorkerUpdateTTL.NewUpdateTTL(a.com).Run); err != nil {
		log.Fatal(err)
	}

	// Будем каждые 1мин пересчитывать различные TTL у городов
	if err := cron.AddFunc("0 * * * * *", cronCitiesUpdateTTL.NewUpdateTTL(a.com).Run); err != nil {
		log.Fatal(err)
	}

	cron.Start()
}

// Инициализация приложения
func (a *App) Init(configuration interface{}) {
	a.configuration = configuration.(*config.Configuration)

	// Инициализация работы с БД
	a.initDB()

	// Инициализация веб-сервера
	a.initWebServer()

	// Инициализация micro
	a.initMicro()

	// Инициализация ядра
	a.pub = corePubs.NewPubs()
	a.base = coreBase.NewBase(a.configuration, a.db, a.service)
	a.com = coreComponents.NewComponents(a.pub, a.base)

	// Инициализация публикаторов-подписчиков
	a.initCityCitiesPubSub()
	a.initTaskWorkerSearchPubSub()
	a.initSubWorkerRegister()
	a.initSubWorkerWorkers()

	// Роутинг запросов
	a.webServer.Get("/api/city/add_city/:name", add_city.NewAddCity(a.com).Handler)

	// Иенициализация компонентов
	a.com.Init()

	// Инициализация планировщика заданий
	a.initCron()
}

// Обработчик приложения
func (a *App) Run() {
	// defer the close till after the main function has finished executing
	defer a.db.Close()

	// Обработчик micro запускаем на отдельной горутине
	go func() {
		// Run server
		if err := a.service.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	// Указываем свой хост и порт и слушаем его
	a.webServer.RunOnAddr(a.configuration.Server.Host + ":" + strconv.Itoa(a.configuration.Server.Port))
}
