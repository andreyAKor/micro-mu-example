package city

import (
	"time"

	natsTopic "nats/topic"

	natsNode "nats/node"
	natsTopicManager "nats/topic/manager"

	protoCityCity "proto/city/city"
	protoTaskWorkerSearch "proto/task/worker/search"

	coreBase "geo.manager/core/base"
	corePubs "geo.manager/core/pubs"

	// DB
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/micro/go-micro/client"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// Список сервисов сети geo, которые использует текущий сервис
const (
	GeoWorker = "geo.worker"
)

// Структура City для работы с городами
type City struct {
	pub  *corePubs.Pubs
	base *coreBase.Base
}

// Конструктор
func NewCity(pub *corePubs.Pubs, base *coreBase.Base) *City {
	// Определяем значения по умолчанию
	return &City{
		pub:  pub,
		base: base,
	}
}

// Поиск станций
func (a *City) GetCity(name string) error {
	log.WithFields(log.Fields{
		"name": name,
	}).Info("Поиск станций")

	// У ворера узнаем список городов
	err := natsTopic.PubRequest(a.pub.PubWorkerCityStations, &protoCityCity.CityEventRequest{
		ResponseTopic: natsTopic.Topic(a.base.Configuration.App.Name, natsTopicManager.CityResponseStations),
		Name:          name,
	})

	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Components.Rpc.City.City.GetCity: PubRequest")

		return err
	}

	return nil
}

// Ставим задачу воркеру для пробивки города
func (a *City) TaskSearchRun(workerUUID string, name string, listWorkerUUID *[]string) error {
	log.WithFields(log.Fields{
		"workerUUID":     workerUUID,
		"name":           name,
		"listWorkerUUID": listWorkerUUID,
	}).Info("Ставим задачу воркеру для пробивки города")

	rsp, err := protoTaskWorkerSearch.NewTaskWorkerSearchRunService(
		GeoWorker,
		natsNode.Client(
			workerUUID,
			a.base.Service.Client(),
		),
	).TaskWorkerSearchRun(
		context.Background(),
		&protoTaskWorkerSearch.TaskWorkerSearchRunRpcRequest{
			ResponseTopic:  natsTopic.Topic(a.base.Configuration.App.Name, natsTopicManager.TaskWorkerResponseSearch),
			LogTopic:       natsTopic.Topic(a.base.Configuration.App.Name, natsTopicManager.TaskWorkerLogSearch),
			Name:           name,
			ListWorkerUUID: *listWorkerUUID,
		},
		client.CallOption(
			client.WithRequestTimeout(time.Duration(time.Second)*3),
		),
	)

	if err != nil {
		log.WithFields(log.Fields{
			"err":            err.Error(),
			"workerUUID":     workerUUID,
			"name":           name,
			"listWorkerUUID": listWorkerUUID,
			"rsp":            rsp,
		}).Error("Components.Rpc.City.City.TaskSearchRun: TaskSearchRun")

		return err
	}

	return nil
}

// Ставим задачу воркеру для остановки пробивки города
func (a *City) TaskWorkerSearchStop(workerUUID string, name string) error {
	log.WithFields(log.Fields{
		"workerUUID": workerUUID,
		"name":       name,
	}).Info("Ставим задачу воркеру для остановки пробивки города")

	rsp, err := protoTaskWorkerSearch.NewTaskWorkerSearchStopService(
		GeoWorker,
		natsNode.Client(
			workerUUID,
			a.base.Service.Client(),
		),
	).TaskWorkerSearchStop(
		context.Background(),
		&protoTaskWorkerSearch.TaskWorkerSearchStopRpcRequest{
			Name: name,
		},
		client.CallOption(
			client.WithRequestTimeout(time.Duration(time.Second)*10),
		),
	)

	if err != nil {
		log.WithFields(log.Fields{
			"err":        err.Error(),
			"workerUUID": workerUUID,
			"name":       name,
			"rsp":        rsp,
		}).Error("Components.Rpc.City.City.TaskWorkerSearchStop: TaskWorkerSearchStop")

		return err
	}

	return nil
}
