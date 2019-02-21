package log

import (
	"time"

	protoTaskWorkerSearch "proto/task/worker/search"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/syncmap"
)

// Аггрегатор лог-данных
type Aggregator struct {
	messages *syncmap.Map
}

// Конструктор
func NewAggregator() *Aggregator {
	log.Debug("Construct NewAggregator")

	return &Aggregator{
		messages: &syncmap.Map{},
	}
}

// Добавляет лог-данные в aггрегатор
func (this *Aggregator) Add(message string) error {
	log.WithFields(log.Fields{
		"message": message,
	}).Debug("Call Add")

	// Получаем сообщение, если оно есть иначе добавляем новое сообщение
	value, loaded := this.messages.LoadOrStore(message, &protoTaskWorkerSearch.TaskWorkerSearchRunEventLog_Messages{
		Message: message,
		Counter: 1,
		Ts:      time.Now().Format(time.RFC3339),
	})

	// Если сообщение уже есть в списке, то обновляем его данные
	if loaded {
		// cast value to correct format
		msg, ok := value.(*protoTaskWorkerSearch.TaskWorkerSearchRunEventLog_Messages)

		if ok {
			msg.Counter++
			msg.Ts = time.Now().Format(time.RFC3339)

			this.messages.Store(message, msg)
		}
	}

	return nil
}

// Возвращает все лог-данные из aггрегатора
func (this *Aggregator) All() []*protoTaskWorkerSearch.TaskWorkerSearchRunEventLog_Messages {
	log.Debug("Call All")

	var messages []*protoTaskWorkerSearch.TaskWorkerSearchRunEventLog_Messages

	this.messages.Range(func(key, value interface{}) bool {
		// cast value to correct format
		msg, ok := value.(*protoTaskWorkerSearch.TaskWorkerSearchRunEventLog_Messages)

		if !ok {
			return false
		}

		messages = append(messages, msg)

		return true
	})

	return messages
}

// Очищает лог-данные в aггрегаторе
func (this *Aggregator) Flush() {
	log.Debug("Call Flush")

	this.messages.Range(func(key, value interface{}) bool {
		this.messages.Delete(key)

		return true
	})
}
