package helpers

import (
	"time"

	coreComponents "geo.manager/core/components"
	stModels "geo.manager/structs/models"

	"github.com/Jeffail/tunny"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

type sendSearchParam struct {
	c  *coreComponents.Components
	ct stModels.Cities
}

// Пулл горутин для работы с search
var poolRunSendSearch = tunny.NewFunc(90, func(param interface{}) interface{} {
	return sendSearch(param.(*sendSearchParam).c, param.(*sendSearchParam).ct)
})

func SendSearch(c *coreComponents.Components, ct stModels.Cities) error {
	log.WithFields(log.Fields{
		"name": ct.Name,
	}).Info("Helpers.SendSearch")

	// Обработчик пулла транзакций запускаем на отдельной горутине
	go func() {
		// Ловим паники
		defer func() {
			if str := recover(); str != nil {
				log.WithFields(log.Fields{
					"err":  str,
					"name": ct.Name,
				}).Error("Helpers.SendSearch: go.recover")
			}
		}()

		// Обработчик транзакции
		if err := poolRunSendSearch.Process(&sendSearchParam{c, ct}); err != nil {
			log.WithFields(log.Fields{
				"err":  err.(error).Error(),
				"name": ct.Name,
			}).Error("Helpers.SendSearch: runSendSearch")
		}
	}()

	return nil
}

func sendSearch(c *coreComponents.Components, ct stModels.Cities) error {
	log.WithFields(log.Fields{
		"name": ct.Name,
	}).Info("Helpers.sendSearch")

	if ct.Status != stModels.CtStatusOpen || (ct.WorkerUUID != nil && len(*ct.WorkerUUID) > 0) {
		return nil
	}

	// Получаем список воркеров
	workersList := c.DbWorker.GetAll(nil)
	if workersList == nil {
		return nil
	}

	// Максималный коэффициент нагрузкаи на воркерами
	maxLoadFactor := float64(0)

	// Выбранная рабочий воркер
	var workingWorker stModels.Workers

	// Перебираем список воркеров и определяем коэффициент нагрузки
	for _, worker := range *workersList {
		// Нулевки игнорим
		if worker.TotalWorkers == nil || uint32(*worker.TotalWorkers) == 0 {
			continue
		}

		// Если израсходовали лимит воркеров у воркеров
		if uint32(*worker.CountWorkers) == uint32(*worker.TotalWorkers) {
			continue
		}

		// Коэффициент нагрузкаи на текущуего воркера
		workerLoadFactor := float64(*worker.TotalWorkers)

		if worker.CountWorkers == nil || uint32(*worker.CountWorkers) > 0 {
			// Реальный коэффициент нагрузки на текущего воркера
			workerLoadFactor = float64(*worker.TotalWorkers) / float64(*worker.CountWorkers)
		}

		// Ищем максимум между предыдущим коэффициентом и текущим
		if maxLoadFactor < workerLoadFactor {
			maxLoadFactor = workerLoadFactor

			if err := copier.Copy(&workingWorker, worker); err != nil {
				log.WithFields(log.Fields{
					"err": err.Error(),
				}).Error("Helpers.SendSearch: copier.Copy")

				return err
			}
		}

	} // for

	// Если нету ни одной свободного воркера
	if len(workingWorker.UUID) == 0 {
		return nil
	}

	log.WithFields(log.Fields{
		"name": ct.Name,
	}).Error("send-search")

	// Список UUID воркеров
	var listWorkerUUID []string

	// Перебираем список воркеров и формируем список воркеров для пробивающего воркера
	for _, worker := range *workersList {
		if worker.UUID != workingWorker.UUID {
			continue
		}

		listWorkerUUID = append(listWorkerUUID, worker.UUID)
	}

	// Ставим задачу воркерe
	if err := c.RpcCityCity.TaskSearchRun(workingWorker.UUID, ct.Name, &listWorkerUUID); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Helpers.SendSearch: RpcCityCity.TaskSearchRun")

		return err
	}

	// Обновляет данные воркера у городов
	c.DbCities.UpdateWorker(ct.Name, workingWorker.UUID, time.Now().Format(time.RFC3339))

	// Пересчитываем счетчики воркера
	freeWorkers := uint32(*workingWorker.FreeWorkers) - 1
	countWorkers := uint32(*workingWorker.CountWorkers) + 1

	// Обновляет данные по воркерам
	c.DbWorker.UpdateWorkers(workingWorker.UUID, time.Now().Format(time.RFC3339), *workingWorker.TotalWorkers, freeWorkers, countWorkers, nil)

	return nil
}
