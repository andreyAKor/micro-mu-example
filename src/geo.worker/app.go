package main

import (
	"os"
	"time"

	core "core"
	natsTopic "nats/topic"
	utilsWorkers "utils/workers"

	// Имена топиков
	natsTopicManager "nats/topic/manager"
	natsTopicWorker "nats/topic/worker"

	protoTaskWorkerSearch "proto/task/worker/search"
	protocWorkerCheckTask "proto/worker/check_task"

	"geo.worker/components"
	"geo.worker/config"
	eventCityCity "geo.worker/event/city/city"
	rpcTaskSearch "geo.worker/rpc/task/task_worker_search"
	rpcWorkerCheckTask "geo.worker/rpc/worker/check_task"

	// Обработчики cron-задач
	cronWorkerRegister "geo.worker/cron/worker/register"
	cronWorkerWorkers "geo.worker/cron/worker/workers"

	// Указываем свои каналы связи
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/transport/nats"

	"github.com/micro/go-micro"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

// Список сервисов сети geo, которые использует текущий сервис
const (
	GeoManager = "geo.manager"
)

// Структура приложения
type App struct {
	core.App

	configuration *config.Configuration
	service       micro.Service
	components    *components.Components
	searchWorkers *utilsWorkers.Workers
}

// Инициализация подписчика для получения списка станций
func (this *App) initSubCityStations() {
	natsTopic.NewSubs(this.configuration.Subscrption.Workers, this.configuration.App, natsTopicWorker.CityRequestStations, this.service, func() interface{} {
		return eventCityCity.NewCity(this.components, this.service)
	})
}

// Инициализация callback для обработки задачи посылки транзакции
func (this *App) initCallTaskSearchRun() {
	protoTaskWorkerSearch.RegisterTaskWorkerSearchRunHandler(
		this.service.Server(),
		rpcTaskSearch.NewTaskWorkerSearchRun(this.components, this.service, this.searchWorkers),
	)
}

// Инициализация callback для остановки обработки задачи посылки транзакции
func (this *App) initCallTaskSearchStop() {
	protoTaskWorkerSearch.RegisterTaskWorkerSearchStopHandler(
		this.service.Server(),
		rpcTaskSearch.NewTaskWorkerSearchStop(this.components, this.service, this.searchWorkers),
	)
}

// Инициализация брокера для запросов Worker.Register
func (this *App) initWorkerRegisterPub() {
	this.components.PubManagerWorkerRegister = natsTopic.NewPub(GeoManager, natsTopicManager.WorkerEventRegister, this.service)
}

// Инициализация брокера для запросов Worker.Workers
func (this *App) initWorkerWorkersPub() {
	this.components.PubManagerWorkerWorkers = natsTopic.NewPub(GeoManager, natsTopicManager.WorkerEventWorkers, this.service)
}

// Регистрируем callback для проверки наличие задачи у воркера
func (this *App) initCallCheckTask() {
	protocWorkerCheckTask.RegisterCheckTaskHandler(
		this.service.Server(),
		rpcWorkerCheckTask.NewCheckTask(this.configuration, this.components, this.searchWorkers),
	)
}

// Инициализация micro
func (this *App) initMicro() {
	// Готовим переменное окружение для micro
	os.Setenv("MICRO_REGISTRY", this.configuration.Micro.Registry)
	os.Setenv("MICRO_REGISTRY_ADDRESS", this.configuration.Micro.RegistryAddress)
	os.Setenv("MICRO_BROKER", this.configuration.Micro.Broker)
	os.Setenv("MICRO_BROKER_ADDRESS", this.configuration.Micro.BrokerAddress)
	os.Setenv("MICRO_TRANSPORT", this.configuration.Micro.Transport)
	os.Setenv("MICRO_TRANSPORT_ADDRESS", this.configuration.Micro.TransportAddress)

	// Экземпляр сервиса
	this.service = micro.NewService(
		micro.Name(this.configuration.App.Name),
		micro.Version(this.configuration.App.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	this.service.Init()
}

// Инициализация планировщика заданий
func (this *App) initCron() {
	// Аля cron на горутинах
	// by https://github.com/robfig/cron + https://stackoverflow.com/questions/16466320/is-there-a-way-to-do-repetitive-tasks-at-intervals-in-golang
	cron := cron.New()

	// Будем каждые 1сек слать регистрационные данные о себе
	if err := cron.AddFunc("* * * * * *", cronWorkerRegister.NewRegister(this.components).Run); err != nil {
		log.Fatal(err)
	}

	// Будем каждые 10сек слать данные о воркерах
	if err := cron.AddFunc("*/10 * * * * *", cronWorkerWorkers.NewWorkers(this.components).Run); err != nil {
		log.Fatal(err)
	}

	cron.Start()
}

// Инициализация приложения
func (this *App) Init(configuration interface{}) {
	this.configuration = configuration.(*config.Configuration)

	// Инициализация micro
	this.initMicro()

	this.searchWorkers = utilsWorkers.NewWorkers("eventTaskSearch", int64(this.configuration.Subscrption.Workers))

	// Инициализация компонентов
	this.components = components.NewComponents(this.configuration, this.service, this.searchWorkers)

	this.initSubCityStations()
	this.initCallTaskSearchRun()
	this.initCallTaskSearchStop()
	this.initWorkerRegisterPub()
	this.initWorkerWorkersPub()
	this.initCallCheckTask()

	this.components.Init()

	// Инициализация планировщика заданий
	this.initCron()
}

// Обработчик приложения
func (this *App) Run() {
	// Run server
	if err := this.service.Run(); err != nil {
		log.Fatal(err)
	}
}
