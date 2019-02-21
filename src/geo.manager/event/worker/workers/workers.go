package workers

import (
	protoWorkerWorkers "proto/worker/workers"

	coreComponents "geo.manager/core/components"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// Сервис
type Workers struct {
	com *coreComponents.Components
}

// Конструктор
func NewWorkers(com *coreComponents.Components) *Workers {
	return &Workers{
		com: com,
	}
}

// Обработчик
func (w *Workers) Process(ctx context.Context, event *protoWorkerWorkers.WorkersEvent) error {
	log.WithFields(log.Fields{
		"event": event,
	}).Info("Received Event.Worker.Workers.Workers.Process")

	// Если воркер не зареган, то регим ее и обновляем ее данные по воркерам
	if worker := w.com.DbWorker.Get(event.Id, nil); worker == nil {
		w.com.DbWorker.Add(event.Id, event.Ts, nil)
	}

	// Обновляет данные по воркерам
	w.com.DbWorker.UpdateWorkers(event.Id, event.Ts, event.TotalWorkers, event.FreeWorkers, event.CountWorkers, nil)

	// Если имеются список воркеров, то перебираем их
	if event.Workers != nil {
		// Перебираем список воркеров
		for _, worker := range event.Workers {
			if cities := w.com.DbCities.Get(worker.Name); cities != nil {
				// Если у города указан воркер отличная от текущей, то у текущей стопорим проверку
				if cities.WorkerUUID != nil && *cities.WorkerUUID != event.Id {
					go func() {
						// Ставим задачу воркер для остановки проверки города
						if err := w.com.RpcCityCity.TaskWorkerSearchStop(event.Id, worker.Name); err != nil {
							log.WithFields(log.Fields{
								"err": err.Error(),
							}).Error("Event.Worker.Workers.Workers.Process: RpcCityCity.TaskWorkerSearchStop")

							//return err
						}
					}()
				} else {
					// Обновляет данные воркера
					w.com.DbCities.UpdateWorker(worker.Name, event.Id, event.Ts)
				}
			} //if
		} // for
	} // if

	return nil
}
