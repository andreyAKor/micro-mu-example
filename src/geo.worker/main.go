package main

import (
	"geo.worker/config"

	core "core"

	log "github.com/sirupsen/logrus"
)

// Точка входа
func main() {
	// Создание службы
	service := core.NewService("geo.worker.conf", new(config.Configuration), new(App))

	// Запуск службы
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
