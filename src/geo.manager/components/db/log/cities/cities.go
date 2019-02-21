package cities

import (
	"time"

	coreBase "geo.manager/core/base"
	corePubs "geo.manager/core/pubs"
	stModels "geo.manager/structs/models"

	// DB
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Структура Cities для работы с логгированием работы с городами
type Cities struct {
	pub  *corePubs.Pubs
	base *coreBase.Base
}

// Конструктор
func NewCities(pub *corePubs.Pubs, base *coreBase.Base) *Cities {
	// Определяем значения по умолчанию
	return &Cities{
		pub:  pub,
		base: base,
	}
}

// Добавляет запись в таблицу логгирования
func (a *Cities) Log(city, type_, message string) {
	a.insert(city, type_, message)
}

// Добавляет запись в таблицу логгирования (расширенный вариант)
func (a *Cities) LogEx(city, type_, message string, counter uint, ts string) {
	a.insertEx(city, type_, message, counter, ts)
}

// Добавляет новую запсиь в таблицу CitiesLog
func (a *Cities) insert(city, type_, message string) {
	a.base.Db.Create(&stModels.CitiesLogs{
		City:    city,
		Type:    type_,
		Message: message,
	})
}

// Добавляет новую запсиь в таблицу CitiesLog (расширенный вариант)
func (a *Cities) insertEx(city, type_, message string, counter uint, ts string) {
	t, _ := time.ParseInLocation(time.RFC3339, ts, time.Now().Location())

	a.base.Db.Create(&stModels.CitiesLogs{
		City:    city,
		Type:    type_,
		Message: message,
		Counter: counter, // Указываем счетчик ошибок

		Model: gorm.Model{
			CreatedAt: t,
		},
	})
}
