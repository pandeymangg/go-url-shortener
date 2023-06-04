package database

import (
	"github.com/pandeymangg/goUrlShortener/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error

	DB, err = gorm.Open(postgres.Open(config.GetConfigString("database.url")), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
