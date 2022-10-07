package repository

import (
	"github.com/zekhoi/learn-golang/entity"
	"gorm.io/gorm"
)

type ShortenRepository interface {
	Create(shorten entity.Shorten) (entity.Shorten, error)
	FindByCode(shortUrl string) (entity.Shorten, error)
	FindAll() ([]entity.Shorten, error)
}

type shortenRepository struct {
	db *gorm.DB
}

func NewShortenRepository(db *gorm.DB) *shortenRepository {
	return &shortenRepository{db}
}

func (r *shortenRepository) Create(shorten entity.Shorten) (entity.Shorten, error) {
	err := r.db.Create(&shorten).Error

	if err != nil {
		return shorten, err
	}

	return shorten, nil
}

func (r *shortenRepository) FindByCode(ShortUrl string) (entity.Shorten, error) {
	var shorten entity.Shorten

	err := r.db.Preload("Shortens").Find(&shorten, "shortUrl = ?", ShortUrl).Error

	if err != nil {
		return shorten, err
	}
	return shorten, nil
}

func (r *shortenRepository) FindAll() ([]entity.Shorten, error) {
	var shortens []entity.Shorten

	err := r.db.Preload("Shortens").Find(&shortens).Error

	if err != nil {
		return shortens, err
	}

	return shortens, nil
}
