package update_ttl

import (
	coreComponents "geo.manager/core/components"

	log "github.com/sirupsen/logrus"
)

// Обновление TTL воркера
type UpdateTTL struct {
	com *coreComponents.Components
}

// Конструктор
func NewUpdateTTL(com *coreComponents.Components) *UpdateTTL {
	return &UpdateTTL{
		com: com,
	}
}

// Обработчик
func (u *UpdateTTL) Run() {
	log.Info("Cron.Worker.UpdateTTL.UpdateTTL.Run: Запуск обработчика планировщика")

	if workersList := u.com.DbWorker.GetAll(nil); workersList != nil {
		for _, worker := range *workersList {
			u.com.DbWorker.UpdateTTL(worker.UUID, nil)
		}
	}
}
