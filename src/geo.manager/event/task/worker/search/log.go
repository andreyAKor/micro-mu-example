package search

import (
	protoTaskWorkerSearch "proto/task/worker/search"

	coreComponents "geo.manager/core/components"
	stModels "geo.manager/structs/models"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// Сервис для данных по Log
type Log struct {
	com *coreComponents.Components
}

// Конструктор
func NewLog(com *coreComponents.Components) *Log {
	return &Log{
		com: com,
	}
}

// Логирование данных search
func (l *Log) Process(ctx context.Context, event *protoTaskWorkerSearch.TaskWorkerSearchRunEventLog) error {
	log.WithFields(log.Fields{
		"event": *event,
	}).Info("Received Event.Task.Worker.Search.Log.Process")

	// Если была ошибка
	if event.Messages != nil {
		// Перебираем список ошибок
		for _, message := range event.Messages {
			log.WithFields(log.Fields{
				"Name":            event.Name,
				"message.Message": message.Message,
				"message.Counter": message.Counter,
				"message.Ts":      message.Ts,
			}).Info("При совершении проверкаи города ошибка")

			// Логгирование ошибки
			l.com.DbLogCities.LogEx(event.Name, stModels.TypeError, message.Message, uint(message.Counter), message.Ts)
		}
	}

	return nil
}
