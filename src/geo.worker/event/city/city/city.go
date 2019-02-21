package accounts

import (
	"time"

	natsTopic "nats/topic"

	protoCityCity "proto/city/city"

	"geo.worker/components"

	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// Сервис для получения списка станций
type City struct {
	components *components.Components
	service    micro.Service
}

// Конструктор
func NewCity(components *components.Components, service micro.Service) *City {
	log.Debug("Construct NewCity")

	return &City{
		components: components,
		service:    service,
	}
}

// Обработчик
func (this *City) Process(ctx context.Context, event *protoCityCity.CityEventRequest) error {
	log.WithFields(log.Fields{
		"event": event,
	}).Info("Received Event.City.City.City.Process")

	// Формируем ответ как запрос на обратное событие
	rsp := &protoCityCity.CityEventResponse{
		Name: event.Name,
	}

	stations := [5]string{"Мытищи", "Пушкино", "Подольск", "Фрязино", "Монино"}

	// Типа что-то делаем и получаем список станций
	rsp.Stations = stations[0:5]

	// Ну а если был ошибка, то пишем ее сюда
	//rsp.Error = err.Error()

	// Временная метка операции
	rsp.Ts = time.Now().Format(time.RFC3339)

	// Ждем 10 сек, типа тут работа кипит во всю))
	time.Sleep(10 * time.Second)

	// Публикатор для ответа на запрос
	return natsTopic.PubResponse(event.ResponseTopic, rsp, this.service)
}
