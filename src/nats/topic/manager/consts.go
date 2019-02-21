// Список имен топиков для работы с сервисом geo.manager

package manager

// Константы топиков-ответов
const (
	// Ответ на получение списка станций у города
	CityResponseStations = "City.Response.Stations"
)

// Константы топиков-запросов по задачам
const (
	// Ответ на обработку задачи пробивки станций у города
	TaskWorkerResponseSearch = "Task.Worker.Response.Search"

	// Ответ с лог-данными на обработку задачи пробивки станций у города
	TaskWorkerLogSearch = "Task.Worker.Log.Search"
)

// Константы топиков-событий
const (
	// Событие на регистрацию воркера
	WorkerEventRegister = "Worker.Event.Register"

	// Событие на данные по воркерам у воркеров
	WorkerEventWorkers = "Worker.Event.Workers"
)
