package register

import (
	protoWorkerRegister "proto/worker/register"

	coreComponents "geo.manager/core/components"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// Сервис
type Register struct {
	com *coreComponents.Components
}

// Конструктор
func NewRegister(com *coreComponents.Components) *Register {
	return &Register{
		com: com,
	}
}

// Обработчик
func (r *Register) Process(ctx context.Context, event *protoWorkerRegister.RegisterEvent) error {
	log.WithFields(log.Fields{
		"event": event,
	}).Info("Received Event.Worker.Register.Register.Process")

	if worker := r.com.DbWorker.Get(event.Id, nil); worker != nil {
		r.com.DbWorker.Update(event.Id, event.Ts, nil)
	} else {
		r.com.DbWorker.Add(event.Id, event.Ts, nil)
	}

	return nil
}
