package components

import (
	utilsWorkers "utils/workers"

	comWorker "geo.worker/components/worker"
	"geo.worker/config"

	"github.com/micro/go-micro"
)

// Структура для работы с компонентами
type Components struct {
	Configuration *config.Configuration
	Service       micro.Service
	SearchWorkers *utilsWorkers.Workers

	Worker *comWorker.Worker

	// Публикаторы
	PubManagerWorkerRegister micro.Publisher
	PubManagerWorkerWorkers  micro.Publisher
}

// Конструктор
func NewComponents(configuration *config.Configuration, service micro.Service, searchWorkers *utilsWorkers.Workers) *Components {
	// Определяем значения по умолчанию
	return &Components{
		Configuration: configuration,
		Service:       service,
		SearchWorkers: searchWorkers,
	}
}

// Инициализация компонентов
func (this *Components) Init() {
	this.Worker = comWorker.NewWorker(this.Configuration, this.Service, this.SearchWorkers, this.PubManagerWorkerRegister, this.PubManagerWorkerWorkers)
}
