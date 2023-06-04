package repository

import (
	"fmt"

	"github.com/pandeymangg/goUrlShortener/models/entity"
	"github.com/pandeymangg/goUrlShortener/utils"
	"gorm.io/gorm"
)

var UrlInstance UrlRepository

type UrlRepository interface {
	CreateShortUrl(url string) (entity.Url, error)
	GetOriginalUrl(shortUrl string) (string, error)
}

type UrlRepositoryImpl struct {
	DB *gorm.DB
}

func NewUrlRepository(db *gorm.DB) UrlRepository {
	if UrlInstance == nil {
		UrlInstance = &UrlRepositoryImpl{DB: db}
	}

	return UrlInstance
}

var SHORT_URL_LENGTH = 6

func (u UrlRepositoryImpl) CreateShortUrl(url string) (entity.Url, error) {

	var urlEntity *entity.Url
	shortUrl := utils.GenerateRandomString(SHORT_URL_LENGTH)
	urlEntity = &entity.Url{ShortUrl: shortUrl}

	u.DB.Where(urlEntity).First(urlEntity)
	for urlEntity.Id != 0 {
		urlEntity.ShortUrl = utils.GenerateRandomString(SHORT_URL_LENGTH)
		u.DB.Where(urlEntity).First(urlEntity)
	}

	urlEntity.Original = url

	u.DB.Create(urlEntity)
	return *urlEntity, nil
}

func (u UrlRepositoryImpl) GetOriginalUrl(shortUrl string) (string, error) {
	urlEntity := &entity.Url{ShortUrl: shortUrl}

	u.DB.Find(urlEntity)

	if urlEntity.Id == 0 {
		return "", fmt.Errorf("no record found for %s", shortUrl)
	}

	return urlEntity.Original, nil
}
