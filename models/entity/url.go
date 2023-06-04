package entity

import "time"

type Url struct {
	Id        int    `gorm:"primaryKey"`
	ShortUrl  string `gorm:"unique"`
	Original  string
	Visits    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Url) TableName() string {
	return "urls"
}
