package check_task

import (
	"time"

	utilsWorkers "utils/workers"

	protocWorkerCheckTask "proto/worker/check_task"

	"geo.worker/components"
	"geo.worker/config"
	rpcTaskSearch "geo.worker/rpc/task/task_worker_search"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// Сервис для запроса на CheckTask
type CheckTask struct {
	configuration *config.Configuration
	components    *components.Components

	// Работа с воркерами
	workers *utilsWorkers.Workers
}

// Конструктор
func NewCheckTask(configuration *config.Configuration, components *components.Components, workers *utilsWorkers.Workers) *CheckTask {
	log.Debug("Construct Rpc.Worker.CheckTask.CheckTask.NewCheckTask")

	return &CheckTask{
		configuration: configuration,
		components:    components,
		workers:       workers,
	}
}

// Обработчик
func (this *CheckTask) CheckTask(ctx context.Context, req *protocWorkerCheckTask.CheckTaskRpcRequest, rsp *protocWorkerCheckTask.CheckTaskRpcResponse) error {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("Received Rpc.Worker.CheckTask.CheckTask.CheckTask")

	// Генерируем ключ воркера
	id := this.workers.GenWorkerId(&rpcTaskSearch.WorkerKey{
		Name: req.Name,
	})

	// Признак пробиваемого города
	rsp.InProgress = true

	// Если ошибка, то воркер отсутсвует
	if _, err := this.workers.GetWorker(id); err != nil {
		// Признак пробиваемого города
		rsp.InProgress = false
	}

	// Временная метка операции
	rsp.Ts = time.Now().Format(time.RFC3339)

	return nil
}
