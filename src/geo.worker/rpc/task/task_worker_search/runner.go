package task_worker_search

import (
	"time"

	natsTopic "nats/topic"

	utilsWorkers "utils/workers"

	protoTaskWorkerSearch "proto/task/worker/search"

	"geo.worker/components"
	ettLog "geo.worker/rpc/task/task_worker_search/log"

	"github.com/Jeffail/tunny"
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
)

// Сервис для пробивки станций
type Runner struct {
	components *components.Components
	service    micro.Service

	// Работа с воркерами
	workers *utilsWorkers.Workers

	// ID текущего воркера
	workerId string

	// Данные события
	Req *protoTaskWorkerSearch.TaskWorkerSearchRunRpcRequest

	// Пулл горутин для работы с пробивками
	poolRunSendSearch *tunny.Pool

	// Аггрегатор лог-данных
	logAggregator *ettLog.Aggregator

	// Канал для обмена сообщением с горутиной обработчика пробивки
	runSendSearchChan chan string
}

// Конструктор
func NewRunner(components *components.Components, service micro.Service, req *protoTaskWorkerSearch.TaskWorkerSearchRunRpcRequest, workers *utilsWorkers.Workers) *Runner {
	log.Debug("Construct NewRunner")

	return &Runner{
		components: components,
		service:    service,
		Req:        req,
		workers:    workers,

		// Аггрегатор лог-данных
		logAggregator: ettLog.NewAggregator(),

		// Канал для обмена сообщением с горутиной обработчика пробивки
		runSendSearchChan: make(chan string),
	}
}

// Обработчик
func (this *Runner) Run() error {
	log.WithFields(log.Fields{
		"req": this.Req,
	}).Debug("Call Event.Task.Runner.Runner.Run")

	// Ловим паники
	defer func() {
		if str := recover(); str != nil {
			log.WithFields(log.Fields{
				"err": str,
				"req": *this.Req,
			}).Error("Error Rpc.Task.TaskWorkerSearch.Runner.Runner.Run: recover")

			// Публикуем лог
			this.publishLog()

			// Дерегистрирует воркера
			_ = this.workers.UnRegister(this.workerId)
		}
	}()

	// Регистрирует нового воркера
	workerId, err := this.workers.Register(
		this.workers.GenWorkerId(WorkerKey{
			Name: this.Req.Name,
		}),
		WorkerData{
			Runner: this,
		},
	)
	if err != nil {
		log.WithFields(log.Fields{
			"err":  err.Error(),
			"name": "eventTaskRunner",
			"req":  *this.Req,
		}).Error("Error Rpc.Task.TaskWorkerSearch.Runner.Runner.Run: workers.Register")

		// Публикуем одну лог-запись
		this.publishOneLog(err.Error())

		return nil
	}

	// ID текущего воркера
	this.workerId = workerId

	log.WithFields(log.Fields{
		"workerId": this.workerId,
	}).Debug("Event.Task.Runner.Runner.Run")

	// Пулл горутин для работы с пробивками
	this.poolRunSendSearch = tunny.NewFunc(6, func(this interface{}) interface{} {
		return this.(*Runner).runSendSearch()
	})

	// При выходе из функции, убиваем пулл горутин для работы с пробивками
	defer this.poolRunSendSearch.Close()

	// Основной цикл
	for {
		// Обработчик пулла пробивок запускаем на отдельной горутине
		go func() {
			// Ловим паники
			defer func() {
				if str := recover(); str != nil {
					log.WithFields(log.Fields{
						"err":  str,
						"name": this.Req.Name,
					}).Error("Error Rpc.Task.TaskWorkerSearch.Runner.Runner.Run: go.recover")
				}
			}()

			// Обработчик пробивки
			if runSendSearchErr := this.poolRunSendSearch.Process(this); runSendSearchErr != nil {
				log.WithFields(log.Fields{
					"err":  runSendSearchErr.(error).Error(),
					"name": this.Req.Name,
				}).Error("Error Rpc.Task.TaskWorkerSearch.Runner.Runner.Run: runSendSearch")

				// Публикуем ответ
				if err := this.publishResponse(runSendSearchErr.(error).Error()); err != nil {
					log.WithFields(log.Fields{
						"err": err.Error(),
					}).Error("Error Rpc.Task.TaskWorkerSearch.Runner.Runner.Run: publishResponse")
				}

				// Стопорим обработку
				this.Stop()
			}

			// Шлем в канал сообщение об продолжении цикла
			this.runSendSearchChan <- "nextFor"
		}()

		// Если в канале сообщение об остановке цикла, то выходим из цикла
		if <-this.runSendSearchChan == "endFor" {
			break
		}
	} // for

	// Публикуем лог
	this.publishLog()

	// Дерегистрирует воркера
	if err := this.workers.UnRegister(this.workerId); err != nil {
		return err
	}

	log.Debug("Event.Task.Runner.Runner.Run: The end")

	return nil
}

// Стопорим обработку
func (this *Runner) Stop() {
	// Шлем в канал сообщение об остановке цикла
	this.runSendSearchChan <- "endFor"
}

// Обработчик пробивки
func (this *Runner) runSendSearch() error {
	log.Debug("Call Event.Task.Runner.Runner.runSendSearch")

	// Ждем 30 сек, типа тут работа кипит))
	time.Sleep(30 * time.Second)

	log.WithFields(log.Fields{
		"name": this.Req.Name,
	}).Debug("Event.Task.Runner.Runner.runSendSearch")

	// Публикуем ответ
	if err := this.publishResponse(""); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Error Rpc.Task.TaskWorkerSearch.Runner.Runner.runSendSearch: publishResponse")

		return err
	}

	// Стопорим обработку
	this.Stop()

	return nil
}

// Публикуем лог
func (this *Runner) publishLog() {
	log.Debug("Call Event.Task.Runner.Runner.publishLog")

	// Формируем лог как запрос на обратное событие
	rsp := &protoTaskWorkerSearch.TaskWorkerSearchRunEventLog{
		Name:     this.Req.Name,
		Messages: this.logAggregator.All(), // Возвращает все лог-данные из aггрегатора
		Ts:       time.Now().Format(time.RFC3339),
	}

	// Очищает лог-данные в aггрегаторе
	this.logAggregator.Flush()

	// Публикатор для ответа на запрос
	_ = natsTopic.PubResponse(this.Req.LogTopic, rsp, this.service)
}

// Публикуем одну лог-запись
func (this *Runner) publishOneLog(message string) {
	log.WithFields(log.Fields{
		"message": message,
	}).Debug("Call Event.Task.Runner.Runner.publishOneLog")

	// Добавляет лог-данные в aггрегатор
	this.logAggregator.Add(message)

	// Публикуем лог-данные
	this.publishLog()
}

// Публикуем ответ
func (this *Runner) publishResponse(error_ string) error {
	log.WithFields(log.Fields{
		"error": error_,
	}).Debug("Call Event.Task.Runner.Runner.publishResponse")

	// Формируем ответ как запрос на обратное событие
	rsp := &protoTaskWorkerSearch.TaskWorkerSearchRunEventResponse{
		Name:  this.Req.Name,
		Error: error_,
		Ts:    time.Now().Format(time.RFC3339),
	}

	// Публикатор для ответа на запрос
	return natsTopic.PubResponse(this.Req.ResponseTopic, rsp, this.service)
}
