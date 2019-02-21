package models

import (
	// DB
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Структура таблицы CitiesLog
type CitiesLogs struct {
	gorm.Model

	City    string // Название
	Type    string // Тип
	Message string // Сообщение
	Counter uint   // Счетчик ошибок
}
