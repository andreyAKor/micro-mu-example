package base

import (
	"geo.manager/config"

	// DB
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/micro/go-micro"
)

type Base struct {
	Configuration *config.Configuration
	Db            *gorm.DB
	Service       micro.Service
}

func NewBase(configuration *config.Configuration, db *gorm.DB, service micro.Service) *Base {
	return &Base{
		Configuration: configuration,
		Db:            db,
		Service:       service,
	}
}
