package pubs

import (
	"github.com/micro/go-micro"
)

type Pubs struct {
	// Публикаторы
	PubWorkerCityStations micro.Publisher
}

func NewPubs() *Pubs {
	return &Pubs{}
}
