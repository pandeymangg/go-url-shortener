package main

import (
	"github.com/pandeymangg/goUrlShortener/api"
	"github.com/pandeymangg/goUrlShortener/config"
	"github.com/pandeymangg/goUrlShortener/database"
	"github.com/pandeymangg/goUrlShortener/repository"
)

func main() {
	config.InitConfig()
	database.InitDatabase()
	repository.NewUrlRepository(database.DB)
	api.InitRouter()
}
