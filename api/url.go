package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pandeymangg/goUrlShortener/business"
	"github.com/pandeymangg/goUrlShortener/models"
)

func shortenUrl(ctx *gin.Context) {
	var request models.ShortenUrlRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := business.ShortenUrl(ctx, request.Url)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"shortUrl": response})
}

func redirectUrl(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")

	response, err := business.RedirectUrl(ctx, shortUrl)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(302, response)
}
