package workers

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"sync/atomic"

	"golang.org/x/sync/syncmap"
)

// Список ошибок
var (
	ErrWorkerRegistered    = errors.New("Worker has registered")
	ErrWorkersOutLimits    = errors.New("Workers out of limits")
	ErrWorkerNotRegistered = errors.New("Worker has not registered")
	ErrWorkersNotFound     = errors.New("Workers not found")
)

// Даннные воркера
type Worker struct {
	Id    string
	Param interface{}
}

// Структура Workers для работы учетом воркеров
type Workers struct {
	workers    *syncmap.Map // Данные воркеров
	counter    int64        // Счетчик количества воркеров
	prefix     string       // Префикс в идентификаторе воркера
	maxСounter int64        // Максимальное количество воркеров
}

// Конструктор
func NewWorkers(prefix string, maxСounter int64) *Workers {
	// Определяем значения по умолчанию
	return &Workers{
		workers:    &syncmap.Map{},
		counter:    0,
		prefix:     prefix,
		maxСounter: maxСounter,
	}
}

// Регистрирует нового воркера
func (w *Workers) Register(id string, param interface{}) (string, error) {
	if w.GetCount() == w.maxСounter {
		return "", ErrWorkersOutLimits
	}

	//id := w.GenWorkerId(param)

	worker := &Worker{
		Id:    id,
		Param: param,
	}

	if _, ok := w.workers.LoadOrStore(id, worker); ok {
		return "", ErrWorkerRegistered
	}

	atomic.AddInt64(&w.counter, 1)

	return id, nil
}

// Дерегистрирует воркера
func (w *Workers) UnRegister(id string) error {
	if _, ok := w.workers.Load(id); !ok {
		return ErrWorkerNotRegistered
	}

	w.workers.Delete(id)

	atomic.AddInt64(&w.counter, -1)

	return nil
}

// Генерирует идентификатор воркера
func (w *Workers) GenWorkerId(param interface{}) string {
	paramStr := fmt.Sprintf("%v", param)

	hasher := md5.New()
	hasher.Write([]byte(paramStr))

	res := hex.EncodeToString(hasher.Sum(nil))

	if len(w.prefix) > 0 {
		res = w.prefix + "-" + res
	}

	return res
}

// Возвращает количество зарегестрированных воркеров
func (w *Workers) GetCount() int64 {
	return atomic.LoadInt64(&w.counter)
}

// Возвращает максимальное количество воркеров
func (w *Workers) GetMaxCount() int64 {
	return w.maxСounter
}

// Возвращает данные зарегестрированного воркера
func (w *Workers) GetWorker(id string) (*Worker, error) {
	worker, ok := w.workers.Load(id)
	if !ok {
		return nil, ErrWorkerNotRegistered
	}

	return worker.(*Worker), nil
}

// Возвращает список данных зарегестрированных воркеров
func (w *Workers) GetWorkers() ([]*Worker, error) {
	var workers []*Worker

	w.workers.Range(func(key, value interface{}) bool {
		// cast value to correct format
		worker, ok := value.(*Worker)

		if !ok {
			return false
		}

		workers = append(workers, worker)

		return true
	})

	if len(workers) == 0 {
		return nil, ErrWorkersNotFound
	}

	return workers, nil
}
