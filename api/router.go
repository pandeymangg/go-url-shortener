package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pandeymangg/goUrlShortener/config"
)

func InitRouter() {
	router := gin.Default()

	addUrlRoutes(router)

	if err := router.Run(":" + config.GetConfigString("server.port")); err != nil {
		panic(err)
	}
}

func addUrlRoutes(router *gin.Engine) {
	url := router.Group("/url")
	{
		url.POST("/shorten", shortenUrl)
		url.GET("/:shortUrl", redirectUrl)
	}
}
