package models

import (
	"time"

	// DB
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Статусы городов
const (
	CtStatusWait       = "WAIT"        // Ожидает проверки
	CtStatusOpen       = "OPEN"        // Открыт
	CtStatusError      = "ERROR"       // Ошибка (статус ошибки на все случаи жизни)
	CtStatusInProgress = "IN_PROGRESS" // В процессе долбежки
)

// Структура таблицы Cities
type Cities struct {
	gorm.Model

	Name string `gorm:"primary_key"` // Название

	Status string // Статус

	WorkerUUID *string    // Уникальный номер воркера
	TsWorker   *time.Time // Время ответа воркера
}
