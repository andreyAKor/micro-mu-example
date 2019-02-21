package topic

import (
	coreConfig "core/config"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	log "github.com/sirupsen/logrus"
)

// Формирует подписчика
func NewSub(app coreConfig.App, topicName string, handler interface{}, service micro.Service) {
	// Регистрируем подписчика
	// register subscriber with queue, each message is delivered to a unique subscriber
	// Работает по принципу одна очередь и много подписчиков
	if err := micro.RegisterSubscriber(Topic(app.Name, topicName), service.Server(), handler, server.SubscriberQueue(Queue(app.Name, topicName))); err != nil {
		log.WithFields(log.Fields{
			"err":       err.Error(),
			"app":       app,
			"topicName": topicName,
		}).Fatal(err)
	}
}

// Формирует подписчика как набор воркеров
func NewSubs(countWorkers int, app coreConfig.App, topicName string, service micro.Service, fn func() interface{}) {
	// Регистрируем несколько подписчиков, чтобы скопом обрабатывали сообщения из очереди
	for i := 0; i < countWorkers; i++ {
		NewSub(app, topicName, fn(), service)
	}
}
