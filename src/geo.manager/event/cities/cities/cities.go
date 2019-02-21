package cities

import (
	protoWorkerCityCity "proto/city/city"

	coreComponents "geo.manager/core/components"
	stModels "geo.manager/structs/models"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// Сервис для данных по Cities
type Cities struct {
	com *coreComponents.Components
}

// Конструктор
func NewCities(com *coreComponents.Components) *Cities {
	return &Cities{
		com: com,
	}
}

// Проверка данных cities
func (a *Cities) Process(ctx context.Context, event *protoWorkerCityCity.CityEventResponse) error {
	log.WithFields(log.Fields{
		"event": event,
	}).Info("Received Event.Cities.Cities.Cities.Process")

	// Если при получении списка станций была ошибка
	if len(event.Error) > 0 {
		log.WithFields(log.Fields{
			"name":  event.Name,
			"error": event.Error,
		}).Info("При получении списка станций была ошибка")

		a.com.DbLogCities.Log(event.Name, stModels.TypeError, event.Error)

		return nil
	}

	// Создаем запись по городу
	a.com.DbCities.Add(event.Name, event.Ts)

	// Меняем ему статус на открытый
	a.com.DbCities.UpdateStatus(event.Name, stModels.CtStatusOpen)

	return nil
}
