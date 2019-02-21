package search

import (
	coreComponents "geo.manager/core/components"
	helpers "geo.manager/helpers"
	stModels "geo.manager/structs/models"

	log "github.com/sirupsen/logrus"
)

// Сервис для данных по Search
type Search struct {
	com        *coreComponents.Components
	lockSearch bool
}

// Конструктор
func NewSearch(com *coreComponents.Components) *Search {
	return &Search{
		com:        com,
		lockSearch: false,
	}
}

// Проверка данных search
func (t *Search) Run() {
	log.Info("Cron.Task.Search.Search.Run: Запуск обработчика планировщика")

	if t.lockSearch == false {
		t.lockSearch = true

		// Проходит по всему списку городов с определенным статусом и выполняет указанную функцию
		if err := t.com.DbCities.WalkCities(stModels.CtStatusOpen, 0, t.walker); err != nil {
			log.WithFields(log.Fields{
				"err": err.Error(),
			}).Error("Cron.Task.Search.Search.Run: DbCities.WalkCities")
		}

		t.lockSearch = false
	}
}

// Прогульщик по записям городов
func (t *Search) walker(city *stModels.Cities) error {
	log.Info("Cron.Task.Search.Search.Walker: Прогульщик по записям городов")

	// Нету смысла ставить воркеру задачу, если у города стоит признак UUID воркера
	if city.WorkerUUID != nil {
		return nil
	}

	if err := helpers.SendSearch(t.com, *city); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Cron.Task.Search.Search.Walker: Helpers.SendSearch")

		return err
	}

	return nil
}
