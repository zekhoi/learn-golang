package entity

import (
	"time"
)

type Shorten struct {
	ID          int    `gorm:"AUTO_INCREMENT;primaryKey"`
	OriginalUrl string `gorm:"unique"`
	ShortUrl    string `gorm:"type:varchar(20);unique"`
	CreatedAt   time.Time
}
