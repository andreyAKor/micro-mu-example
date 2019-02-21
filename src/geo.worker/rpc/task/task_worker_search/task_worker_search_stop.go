package task_worker_search

import (
	"time"

	utilsWorkers "utils/workers"

	protoTaskWorkerSearch "proto/task/worker/search"

	"geo.worker/components"

	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// Сервис для остановки пробивки станции
type TaskWorkerSearchStop struct {
	components *components.Components
	service    micro.Service

	// Работа с воркерами
	workers *utilsWorkers.Workers
}

// Конструктор
func NewTaskWorkerSearchStop(components *components.Components, service micro.Service, workers *utilsWorkers.Workers) *TaskWorkerSearchStop {
	log.Debug("Construct NewTaskWorkerSearchStop")

	return &TaskWorkerSearchStop{
		components: components,
		service:    service,
		workers:    workers,
	}
}

// Обработчик
func (this *TaskWorkerSearchStop) TaskWorkerSearchStop(ctx context.Context, req *protoTaskWorkerSearch.TaskWorkerSearchStopRpcRequest, rsp *protoTaskWorkerSearch.TaskWorkerSearchStopRpcResponse) error {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("Received Rpc.Task.TaskWorkerSearchStop.TaskWorkerSearchStop.TaskWorkerSearchStop")

	// Запускаем обработчик в горутине, чтобы не блокировать поток обработки очередей
	worker, err := this.workers.GetWorker(this.workers.GenWorkerId(WorkerKey{
		Name: req.Name,
	}))

	if err != nil {
		rsp.Error = err.Error()
	} else {
		worker.Param.(WorkerData).Runner.Stop()
	}

	// Временная метка операции
	rsp.Ts = time.Now().Format(time.RFC3339)

	return nil
}
