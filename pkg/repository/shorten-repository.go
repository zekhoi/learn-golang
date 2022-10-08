package repository

import (
	"github.com/zekhoi/learn-golang/pkg/entity"
	"gorm.io/gorm"
)

type ShortenRepository interface {
	Create(shorten entity.Shorten) (entity.Shorten, error)
	FindByCode(shortUrl string) (entity.Shorten, error)
	FindByOriginal(originalUrl string) (entity.Shorten, error)
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

func (r *shortenRepository) FindByCode(shortUrl string) (entity.Shorten, error) {
	var shorten entity.Shorten

	err := r.db.Where("short_url = ?", shortUrl).First(&shorten).Error

	if err != nil {
		return shorten, err
	}

	return shorten, nil
}

func (r *shortenRepository) FindAll() ([]entity.Shorten, error) {
	var shortens []entity.Shorten

	err := r.db.Find(&shortens).Error

	if err != nil {
		return shortens, err
	}

	return shortens, nil
}

func (r *shortenRepository) FindByOriginal(originalUrl string) (entity.Shorten, error) {
	var shorten entity.Shorten

	if err := r.db.Where("original_url = ?", originalUrl).First(&shorten).Error; err == nil {
		return shorten, nil
	}

	return shorten, nil
}
