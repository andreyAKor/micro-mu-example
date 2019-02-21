package task_worker_search

import (
	"time"

	natsNode "nats/node"
	utilsService "utils/service"
	utilsWorkers "utils/workers"

	protoTaskWorkerSearch "proto/task/worker/search"
	protocWorkerCheckTask "proto/worker/check_task"

	"geo.worker/components"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// Список сервисов сети geo, которые использует текущий сервис
const (
	GeoWorker = "geo.worker"
)

// Сервис для пробивки станции
type TaskWorkerSearchRun struct {
	components *components.Components
	service    micro.Service

	// Работа с воркерами
	workers *utilsWorkers.Workers
}

// Конструктор
func NewTaskWorkerSearchRun(components *components.Components, service micro.Service, workers *utilsWorkers.Workers) *TaskWorkerSearchRun {
	log.Debug("Construct NewTaskWorkerSearchRun")

	return &TaskWorkerSearchRun{
		components: components,
		service:    service,
		workers:    workers,
	}
}

// Обработчик
func (this *TaskWorkerSearchRun) TaskWorkerSearchRun(ctx context.Context, req *protoTaskWorkerSearch.TaskWorkerSearchRunRpcRequest, rsp *protoTaskWorkerSearch.TaskWorkerSearchRunRpcResponse) error {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("Received Rpc.Task.TaskWorkerSearchRun.TaskWorkerSearchRun.TaskWorkerSearchRun")

	// Запускаем обработчик в горутине, чтобы не блокировать поток обработки очередей
	go func() {
		// Имеется ли задача в работе
		isThereTaskInWork := false

		// Если список воркеров имеется, то опрашиваем их на наличие задачи
		if len(req.ListWorkerUUID) > 0 {
			for _, workerUUID := range req.ListWorkerUUID {
				// Себя пропускаем
				if utilsService.Id(this.service) == workerUUID {
					continue
				}

				// Проверяем у соседнего воркера наличие текущей задачи
				rsp, err := protocWorkerCheckTask.NewCheckTaskService(
					GeoWorker,
					natsNode.Client(
						workerUUID,
						this.service.Client(),
					),
				).CheckTask(
					context.Background(),
					&protocWorkerCheckTask.CheckTaskRpcRequest{
						Name: req.Name,
					},
					client.CallOption(
						client.WithRequestTimeout(time.Duration(time.Second)*30),
					),
				)

				if err != nil {
					log.WithFields(log.Fields{
						"err":        err.Error(),
						"workerUUID": workerUUID,
						"req":        *req,
					}).Error("Error Rpc.Task.TaskWorkerSearchRun.TaskWorkerSearchRun.TaskWorkerSearchRun: CheckTask")

					continue
				}

				// Если задача имеется в работе
				if rsp.InProgress == true && !isThereTaskInWork {
					isThereTaskInWork = true
				}
			} // for
		} // if

		// Если задача у других воркеров в работе отсутсвует, то берем текущую задачу мы
		if !isThereTaskInWork {
			// Сервис для пробивки станций
			runner := NewRunner(this.components, this.service, req, this.workers)

			if err := runner.Run(); err != nil {
				log.WithFields(log.Fields{
					"err": err.Error(),
					"req": *req,
				}).Error("Error Rpc.Task.TaskWorkerSearchRun.TaskWorkerSearchRun.TaskWorkerSearchRun: Run")

				//return err
			}
		}
	}()

	log.Debug("Rpc.Task.TaskWorkerSearchRun.TaskWorkerSearchRun.TaskWorkerSearchRun: The end")

	// Временная метка операции
	rsp.Ts = time.Now().Format(time.RFC3339)

	return nil
}
