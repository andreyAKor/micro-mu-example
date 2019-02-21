package components

import (
	coreBase "geo.manager/core/base"
	corePubs "geo.manager/core/pubs"

	comDbCities "geo.manager/components/db/cities"
	comDbLogCities "geo.manager/components/db/log/cities"
	comDbWorker "geo.manager/components/db/worker"
	comRpcCityCity "geo.manager/components/rpc/city/city"
)

// Структура для работы с компонентами
type Components struct {
	Pub  *corePubs.Pubs
	Base *coreBase.Base

	// Компоненты DB
	DbLogCities *comDbLogCities.Cities
	DbCities    *comDbCities.Cities
	DbWorker    *comDbWorker.Worker

	// Компоненты RPC
	RpcCityCity *comRpcCityCity.City
}

// Конструктор
func NewComponents(pub *corePubs.Pubs, base *coreBase.Base) *Components {
	// Определяем значения по умолчанию
	return &Components{
		Pub:  pub,
		Base: base,
	}
}

// Инициализация компонентов
func (c *Components) Init() {
	// Города
	c.DbCities = comDbCities.NewCities(c.Pub, c.Base)
	c.RpcCityCity = comRpcCityCity.NewCity(c.Pub, c.Base)
	c.DbLogCities = comDbLogCities.NewCities(c.Pub, c.Base)

	// Прочее
	c.DbWorker = comDbWorker.NewWorker(c.Pub, c.Base)
}
