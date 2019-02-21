package worker

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

// Структура Worker для работы с воркерами
type Worker struct {
	pub  *corePubs.Pubs
	base *coreBase.Base
}

// Конструктор
func NewWorker(pub *corePubs.Pubs, base *coreBase.Base) *Worker {
	// Определяем значения по умолчанию
	return &Worker{
		pub:  pub,
		base: base,
	}
}

// Возвращает указатель на инстанс работы с базой
func (r *Worker) getDb(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}

	return r.base.Db
}

// Добавляет запись в таблицу Worker
func (r *Worker) Add(uuid, ts string, tx *gorm.DB) {
	log.WithFields(log.Fields{
		"uuid": uuid,
		"ts":   ts,
	}).Info("Регистрируем воркера")

	if worker := r.Get(uuid, tx); worker == nil {
		t, _ := time.ParseInLocation(time.RFC3339, ts, time.Now().Location())

		r.getDb(tx).Create(&stModels.Workers{
			UUID:       uuid,
			TsRegister: &t,
		})
	}
}

// Удаляет запись из таблицы Worker
func (r *Worker) Delete(uuid string, tx *gorm.DB) {
	r.getDb(tx).Unscoped().Where(&stModels.Workers{
		UUID: uuid,
	}).Delete(stModels.Workers{})
}

// Возвращает весь список зарегистрированных воркеров
func (r *Worker) GetAll(tx *gorm.DB) *[]stModels.Workers {
	var workersList []stModels.Workers

	r.getDb(tx).Find(&workersList, stModels.Workers{})

	return &workersList
}

// Обновляет данные
func (r *Worker) Update(uuid, ts string, tx *gorm.DB) {
	if worker := r.Get(uuid, tx); worker != nil {
		t, _ := time.ParseInLocation(time.RFC3339, ts, time.Now().Location())

		ttl := uint(0)

		if worker.TsRegister != nil {
			// Считаем TTL = текущее системное время - последний TS из базы
			ttl = uint(time.Now().Sub(*worker.TsRegister).Seconds())
		}

		if ttl < 0 {
			ttl = uint(0)
		}

		r.getDb(tx).Model(&worker).Updates(stModels.Workers{
			TsRegister: &t,
			TTL:        &ttl,
		})
	}
}

// Обновляет данные TTL
func (r *Worker) UpdateTTL(uuid string, tx *gorm.DB) {
	if worker := r.Get(uuid, tx); worker != nil {
		ttl := uint(0)

		if worker.TsRegister != nil {
			// Считаем TTL = текущее системное время - последний TS из базы
			ttl = uint(time.Now().Sub(*worker.TsRegister).Seconds())
		}

		if ttl < 0 {
			ttl = uint(0)
		}

		r.getDb(tx).Model(&worker).Updates(stModels.Workers{
			TTL: &ttl,
		})
	}
}

// Обновляет данные по воркерам
func (r *Worker) UpdateWorkers(uuid, ts string, totalWorkers, freeWorkers, countWorkers uint32, tx *gorm.DB) {
	if worker := r.Get(uuid, tx); worker != nil {
		t, _ := time.ParseInLocation(time.RFC3339, ts, time.Now().Location())

		ttl := uint(0)

		if worker.TsRegister != nil {
			// Считаем TTL = текущее системное время - последний TS из базы
			ttl = uint(time.Now().Sub(*worker.TsRegister).Seconds())
		}

		if ttl < 0 {
			ttl = uint(0)
		}

		r.getDb(tx).Model(&worker).Updates(stModels.Workers{
			TsRegister:   &t,
			TTL:          &ttl,
			TotalWorkers: &totalWorkers,
			FreeWorkers:  &freeWorkers,
			CountWorkers: &countWorkers,
		})
	}
}

// Получает значение из Worker зная его uuid
func (r *Worker) Get(uuid string, tx *gorm.DB) *stModels.Workers {
	worker := stModels.Workers{}

	res := r.getDb(tx).First(&worker, stModels.Workers{
		UUID: uuid,
	})

	if res.RecordNotFound() {
		return nil
	}

	return &worker
}
