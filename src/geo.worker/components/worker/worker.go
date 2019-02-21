package worker

import (
	"time"

	natsTopic "nats/topic"
	utilsService "utils/service"
	utilsWorkers "utils/workers"

	protoWorkerRegister "proto/worker/register"

	"geo.worker/config"

	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
)

// Структура
type Worker struct {
	configuration            *config.Configuration
	service                  micro.Service
	searchWorkers            *utilsWorkers.Workers
	pubManagerWorkerRegister micro.Publisher
	pubManagerWorkerWorkers  micro.Publisher
}

// Конструктор
func NewWorker(configuration *config.Configuration, service micro.Service, searchWorkers *utilsWorkers.Workers, pubManagerWorkerRegister, pubManagerWorkerWorkers micro.Publisher) *Worker {
	// Определяем значения по умолчанию
	return &Worker{
		configuration:            configuration,
		service:                  service,
		searchWorkers:            searchWorkers,
		pubManagerWorkerRegister: pubManagerWorkerRegister,
		pubManagerWorkerWorkers:  pubManagerWorkerWorkers,
	}
}

// Посылаем регистрационное событие в менеждер
func (this *Worker) Register() error {
	log.WithFields(log.Fields{
		"uuid": utilsService.Id(this.service),
	}).Debug("Call Components.Worker.Worker.Register")

	err := natsTopic.PubRequest(this.pubManagerWorkerRegister, &protoWorkerRegister.RegisterEvent{
		// Уникальный ID воркера
		Id: utilsService.Id(this.service),

		// Временная метка операции
		Ts: time.Now().Format(time.RFC3339),
	})

	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Components.Worker.Worker.Register: PubRequest")

		return err
	}

	return nil
}
