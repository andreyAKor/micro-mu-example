package search

import (
	protoTaskWorkerSearch "proto/task/worker/search"

	coreComponents "geo.manager/core/components"
	stModels "geo.manager/structs/models"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// Сервис для данных по Response
type Response struct {
	com *coreComponents.Components
}

// Конструктор
func NewResponse(com *coreComponents.Components) *Response {
	return &Response{
		com: com,
	}
}

// Проверка данных search
func (r *Response) Process(ctx context.Context, event *protoTaskWorkerSearch.TaskWorkerSearchRunEventResponse) error {
	log.WithFields(log.Fields{
		"event": event,
	}).Info("Received Event.Task.Worker.Search.Response.Process")

	// Если была ошибка
	if len(event.Error) > 0 {
		log.WithFields(log.Fields{
			"Name":  event.Name,
			"error": event.Error,
			"ts":    event.Ts,
		}).Info("При совершении пробивки города была ошибка")

		// Логгирование ошибки
		r.com.DbLogCities.LogEx(event.Name, stModels.TypeError, event.Error, 1, event.Ts)

		return nil
	}

	// Тут что-то прикольное сработало, а что пока ХЗ)
	log.WithFields(log.Fields{
		"Name": event.Name,
	}).Error("Урааааа!!!!")

	// Логгирование инфу
	r.com.DbLogCities.LogEx(event.Name, stModels.TypeInfo, "Урааааа!!!!", 1, event.Ts)

	return nil
}
