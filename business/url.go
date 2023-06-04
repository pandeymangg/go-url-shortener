package business

import (
	"github.com/gin-gonic/gin"
	"github.com/pandeymangg/goUrlShortener/models"
	"github.com/pandeymangg/goUrlShortener/repository"
)

func ShortenUrl(ctx *gin.Context, url string) (models.ShortenUrlResponse, error) {
	shortUrl, err := repository.UrlInstance.CreateShortUrl(url)

	if err != nil {
		return models.ShortenUrlResponse{}, err
	} else {
		return models.ShortenUrlResponse{ShortUrl: shortUrl.ShortUrl}, nil
	}
}

func RedirectUrl(ctx *gin.Context, shortUrl string) (string, error) {
	return repository.UrlInstance.GetOriginalUrl(shortUrl)
}
