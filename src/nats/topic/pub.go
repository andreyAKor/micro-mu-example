package topic

import (
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// Формирует публикатора
func NewPub(serviceName, topicName string, service micro.Service) micro.Publisher {
	return micro.NewPublisher(Topic(serviceName, topicName), service.Client())
}

// Публикатор для ответа на запрос
func PubResponse(responseTopicName string, rsp interface{}, service micro.Service) error {
	log.WithFields(log.Fields{
		"responseTopicName": responseTopicName,
		"rsp":               rsp,
	}).Info("Publish Nats.Topic.PubResponse")

	// Создаем публикатора
	pub := micro.NewPublisher(responseTopicName, service.Client())

	// Публикуем сообщение
	if err := pub.Publish(context.Background(), rsp); err != nil {
		log.WithFields(log.Fields{
			"err":               err.Error(),
			"responseTopicName": responseTopicName,
			"rsp":               rsp,
		}).Error("Nats.Topic.PubResponse: NewPublisher.Publish")

		return err
	}

	return nil
}

// Публикатор для ответа на запрос
func PubRequest(pub micro.Publisher, req interface{}) error {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("Publish Nats.Topic.PubRequest")

	// Публикуем сообщение
	if err := pub.Publish(context.Background(), req); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"req": req,
		}).Error("Nats.Topic.PubRequest: NewPublisher.Publish")

		return err
	}

	return nil
}

// Формирует имя топика
func Topic(serviceName, topicName string) string {
	return serviceName + ".topic." + topicName
}

// Формирует имя очереди
func Queue(serviceName, topicName string) string {
	return serviceName + ".queue." + topicName
}

// Формирует уникальную часть имени топика
func Unique(id, topicName string) string {
	return id + "." + topicName
}
