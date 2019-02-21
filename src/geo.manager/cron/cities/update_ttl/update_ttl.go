package update_ttl

import (
	"time"

	coreComponents "geo.manager/core/components"
	stModels "geo.manager/structs/models"

	log "github.com/sirupsen/logrus"
)

// Обновление TTL воркеров
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
	log.Info("Cron.Cities.UpdateTTL.UpdateTTL.Run: Запуск обработчика планировщика")

	if workersList := u.com.DbCities.GetAll(); workersList != nil {
		// TTL воркеров у городов
		workerTTL := uint(0)

		for _, workers := range *workersList {
			// Считаем данные TTL воркеров у городов
			if workers.TsWorker != nil {
				// Считаем TTL = текущее системное время - последний TS из базы
				workerTTL = uint(time.Now().Sub(*workers.TsWorker).Seconds())
			}

			if workerTTL < 0 {
				workerTTL = uint(0)
			}

			// Если TTL воркеров у городов больше 60 сек, то значит, что город воркером не пробивается и удаляем его из города
			if workerTTL > 60 {
				// Удаляет данные воркера у города
				u.com.DbCities.DeleteWorker(workers.Name)

				// Ставим статус, что город открыт для пробивки
				u.com.DbCities.UpdateStatus(workers.Name, stModels.CtStatusOpen)
			}
		} // for
	} // if
}
