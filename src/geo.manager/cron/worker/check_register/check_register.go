package check_register

import (
	coreComponents "geo.manager/core/components"
	stModels "geo.manager/structs/models"

	log "github.com/sirupsen/logrus"
)

// Проверка актуальности зарегистрированных воркеров по их TTL
type CheckRegister struct {
	com *coreComponents.Components
}

// Конструктор
func NewCheckRegister(com *coreComponents.Components) *CheckRegister {
	return &CheckRegister{
		com: com,
	}
}

// Обработчик
func (c *CheckRegister) Run() {
	log.Info("Cron.Worker.CheckRegister.CheckRegister.Run: Запуск обработчика планировщика")

	if workersList := c.com.DbWorker.GetAll(nil); workersList != nil {
		for _, worker := range *workersList {
			// Если TTL воркера больше 60 сек, то значит, что воркер дохлый и удаляем ее из регистрации
			// Также удаляем упоминание этого воркера из городов
			if worker.TTL != nil && *worker.TTL > 60 {
				// Удаляет запись из таблицы Workers
				c.com.DbWorker.Delete(worker.UUID, nil)

				// Возвращает список городов с UUID определенного воркера
				if citiesList := c.com.DbCities.GetCitiesListByWorkerUUID(worker.UUID); citiesList != nil {
					// Перебираем список городов из базы
					for _, city := range *citiesList {
						// Удаляет данные воркера у города
						c.com.DbCities.DeleteWorker(city.Name)

						// Ставим статус, что город открыт для проверки
						c.com.DbCities.UpdateStatus(city.Name, stModels.CtStatusOpen)
					} // for
				} // if
			} // if
		} // for
	} // if
}
