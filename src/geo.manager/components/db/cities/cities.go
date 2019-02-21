package cities

import (
	"time"

	coreBase "geo.manager/core/base"
	corePubs "geo.manager/core/pubs"
	stModels "geo.manager/structs/models"

	// DB
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	log "github.com/sirupsen/logrus"
)

// Структура Cities для работы с городами
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

// Добавляет запись в таблицу Cities
func (a *Cities) Add(name, ts string) *stModels.Cities {
	log.WithFields(log.Fields{
		"name": name,
		"ts":   ts,
	}).Info("Добавляем новый город")

	cities := a.Get(name)
	if cities == nil {
		t, _ := time.ParseInLocation(time.RFC3339, ts, time.Now().Location())

		cities = &stModels.Cities{
			Name:   name,
			Status: stModels.CtStatusWait,

			Model: gorm.Model{
				CreatedAt: t,
			},
		}

		a.base.Db.Create(cities)
	}

	return cities
}

// Удаляет запись из таблицы Cities по name
func (a *Cities) Delete(name string) {
	a.base.Db.Unscoped().Where(&stModels.Cities{
		Name: name,
	}).Delete(stModels.Cities{})
}

// Обновляет данные воркера у города
func (a *Cities) UpdateWorker(name, workerUUID, tsWorker string) {
	// Получает значение из Cities зная его name
	t, _ := time.ParseInLocation(time.RFC3339, tsWorker, time.Now().Location())

	res := a.base.Db.Model(stModels.Cities{}).Where(stModels.Cities{
		Name: name,
	}).Updates(stModels.Cities{
		WorkerUUID: &workerUUID,
		Status:     stModels.CtStatusInProgress,
		TsWorker:   &t,
	})

	if res.Error != nil {
		log.WithFields(log.Fields{
			"err": res.Error.Error(),
		}).Error("Components.Db.Cities.Cities.UpdateWorker")
	}
}

// Удаляет данные воркера у города
func (a *Cities) DeleteWorker(name string) {
	if cities := a.Get(name); cities != nil {
		cities.WorkerUUID = nil
		cities.TsWorker = nil

		a.base.Db.Save(&cities)
	}
}

// Возвращает весь список городов
func (a *Cities) GetAll() *[]stModels.Cities {
	var citiesList []stModels.Cities

	a.base.Db.Find(&citiesList, stModels.Cities{})

	return &citiesList
}

// Проходит по всему списку городов с определенным статусом и выполняет указанную функцию
func (a *Cities) WalkCities(status string, limit int, fn func(city *stModels.Cities) error) error {
	var citiesList []stModels.Cities

	q := a.base.Db.Order("created_at ASC").Find(&citiesList, stModels.Cities{
		Status: status,
	})

	if limit > 0 {
		q.Limit(limit)
	}

	// Перебираем полученные записи
	for _, cities := range citiesList {
		// Выполняем переданную функцию
		if err := fn(&cities); err != nil {
			return err
		}
	}

	return nil
}

// Возвращает список городов с UUID определенного воркера
func (a *Cities) GetCitiesListByWorkerUUID(workerUUID string) *[]stModels.Cities {
	var citiesList []stModels.Cities

	a.base.Db.Find(&citiesList, stModels.Cities{
		WorkerUUID: &workerUUID,
	})

	return &citiesList
}

// Получает значение из Cities зная его name
func (a *Cities) Get(name string) *stModels.Cities {
	cities := stModels.Cities{}

	res := a.base.Db.First(&cities, stModels.Cities{
		Name: name,
	})

	if res.RecordNotFound() {
		return nil
	}

	return &cities
}

// Обновляет статус города
func (a *Cities) UpdateStatus(name, status string) {
	if cities := a.Get(name); cities != nil && cities.Status != status {
		res := a.base.Db.Model(&cities).Updates(stModels.Cities{
			Status: status,
		})

		if res.Error != nil {
			log.WithFields(log.Fields{
				"name":   name,
				"status": status,
				"err":    res.Error.Error(),
			}).Error("Components.Db.Cities.Cities.UpdateStatus")
		}
	}
}
