package core

import (
	"core/config"

	"github.com/kardianos/service"
	log "github.com/sirupsen/logrus"
)

// Структура программы
type Program struct {
	exit              chan struct{}
	service           service.Service
	app               App
	coreConfiguration *config.Configuration
}

// Обработчик старта сервиса
func (this *Program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	log.Info("Start")

	go this.run()
	return nil
}

// Обработчик программы сервиса
func (this *Program) run() {
	log.Info("Runnig ", this.coreConfiguration.App.DisplayName)

	defer func() {
		// Смотрим наличие менеджеров сервисов в ОС
		if service.ChosenSystem() != nil {
			if service.Interactive() {
				this.Stop(this.service)
			} else {
				this.service.Stop()
			}
		}
	}()

	// Запускаем обработчик приложения
	this.app.Run()
}

// Обработчик остановки сервиса
func (this *Program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	log.Info("Stop")

	return nil
}
