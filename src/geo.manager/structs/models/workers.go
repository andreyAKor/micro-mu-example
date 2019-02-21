package models

import (
	"time"

	// DB
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Структура таблицы Workers
type Workers struct {
	gorm.Model

	UUID       string     // Уникальный номер воркера
	TsRegister *time.Time // Время регистрации
	TTL        *uint      // TTL

	TotalWorkers *uint32 // Общее количество воркеров у воркера
	FreeWorkers  *uint32 // Количество свободных воркеров у воркера
	CountWorkers *uint32 // Количество занятых воркеров у воркера
}
