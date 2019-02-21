package register

import (
	"geo.worker/components"

	log "github.com/sirupsen/logrus"
)

// Сервис для данных по Worker.Register
type Register struct {
	Components *components.Components
}

// Конструктор
func NewRegister(components *components.Components) *Register {
	return &Register{
		Components: components,
	}
}

// Посылаем регистрационное событие в менеждер
func (this *Register) Run() {
	log.Info("Cron.Worker.Register.Register.Run: Запуск обработчика планировщика")

	// Посылаем регистрационное событие в менеждер
	if err := this.Components.Worker.Register(); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Cron.Worker.Register.Register.Run: Worker.Register")
	}
}
