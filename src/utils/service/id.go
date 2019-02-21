package service

import (
	"github.com/micro/go-micro"
)

// Уникальный ID сервиса
func Id(service micro.Service) string {
	return service.Server().Options().Name + "-" + service.Server().Options().Id
}
