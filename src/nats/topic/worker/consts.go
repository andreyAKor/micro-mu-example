// Список имен топиков для работы с сервисом geo.worker

package worker

// Константы топиков-запросов
const (
	// Запрос на получение списка станций у города
	CityRequestStations = "City.Request.Stations"

	// Запрос на обработку задачи пробивки станций у города
	TaskRequestSearch = "Task.Request.Search"
)
