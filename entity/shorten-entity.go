package entity

import "time"

type Shorten struct {
	ID          int `gorm:"primaryKey"`
	OriginalUrl string
	ShortUrl    string
	CustomUrl   string
	CreatedAt   time.Time
}
