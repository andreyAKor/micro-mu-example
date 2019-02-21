package workers

import (
	"time"

	natsTopic "nats/topic"
	utilsService "utils/service"

	protoWorkerWorkers "proto/worker/workers"

	"geo.worker/components"
	rpcTaskSearch "geo.worker/rpc/task/task_worker_search"

	log "github.com/sirupsen/logrus"
)

// Сервис для данных по Worker.Workers
type Workers struct {
	Components *components.Components
}

// Конструктор
func NewWorkers(components *components.Components) *Workers {
	return &Workers{
		Components: components,
	}
}

// Посылаем событие с данными воркеров в менеждер
func (this *Workers) Run() {
	log.Info("Cron.Worker.Workers.Workers.Run: Запуск обработчика планировщика")

	// Посылаем событие с данными воркеров в менеждер
	if err := this.workers(); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Cron.Worker.Workers.Workers.Run: Worker.Workers")
	}
}

// Посылаем событие с данными воркеров в менеждер
func (this *Workers) workers() error {
	log.WithFields(log.Fields{
		"uuid": utilsService.Id(this.Components.Service),
	}).Debug("Call Cron.Worker.Workers.Workers.Run")

	totalWorkers := this.Components.SearchWorkers.GetMaxCount()
	countWorkers := this.Components.SearchWorkers.GetCount()
	freeWorkers := totalWorkers - countWorkers

	// Костыль, на случай, если количество свободных воркеров больше общего допустимого количества воркеров
	if freeWorkers > totalWorkers {
		freeWorkers = totalWorkers
	}

	// Список воркеров для события
	var workers []*protoWorkerWorkers.WorkerData

	// Возвращает список данных зарегестрированных воркеров
	if workersList, err := this.Components.SearchWorkers.GetWorkers(); err == nil {
		// Перебираем список воркеров
		for _, worker := range workersList {
			workers = append(workers, &protoWorkerWorkers.WorkerData{
				// Название
				Name: worker.Param.(rpcTaskSearch.WorkerData).Runner.Req.Name,
			})
		}
	}

	err := natsTopic.PubRequest(this.Components.PubManagerWorkerWorkers, &protoWorkerWorkers.WorkersEvent{
		// Уникальный ID воркера
		Id: utilsService.Id(this.Components.Service),

		// Общее количество воркеров
		TotalWorkers: uint32(totalWorkers),

		// Количество свободных воркеров
		FreeWorkers: uint32(freeWorkers),

		// Количество занятых воркеров
		CountWorkers: uint32(countWorkers),

		// Список данных занятых воркеров
		Workers: workers,

		// Временная метка операции
		Ts: time.Now().Format(time.RFC3339),
	})

	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Cron.Worker.Workers.Workers.Run: PubRequest")

		return err
	}

	return nil
}
