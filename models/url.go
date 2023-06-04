package models

type ShortenUrlRequest struct {
	Url string `json:"url"`
}
type ShortenUrlResponse struct {
	ShortUrl string `json:"shortUrl"`
}
