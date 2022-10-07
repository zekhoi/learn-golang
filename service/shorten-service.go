package service

import (
	"github.com/zekhoi/learn-golang/entity"
	"github.com/zekhoi/learn-golang/repository"
	"github.com/zekhoi/learn-golang/request"
)

type ShortenService interface {
	CreateShorten(request request.CreateShortenRequest) (entity.Shorten, error)
	GetShortenByCode(request request.GetShortenRequest) (entity.Shorten, error)
	GetShortens() ([]entity.Shorten, error)
}

type shortenService struct {
	repository repository.ShortenRepository
}

func NewShortenService(repository repository.ShortenRepository) *shortenService {
	return &shortenService{repository}
}

func (s *shortenService) GetShortens() ([]entity.Shorten, error) {
	shortens, err := s.repository.FindAll()

	if err != nil {
		return shortens, nil
	}

	return shortens, nil
}

func (s *shortenService) GetShortenByCode(request request.GetShortenRequest) (entity.Shorten, error) {
	shorten, err := s.repository.FindByCode(request.ShortUrl)
	// fmt.Println(order)
	if err != nil {
		return shorten, err
	}

	return shorten, nil
}

func (s *shortenService) CreateShorten(request request.CreateShortenRequest) (entity.Shorten, error) {
	shorten := entity.Shorten{}
	shorten.OriginalUrl = request.OriginalUrl
	shorten.CustomUrl = request.CustomUrl

	newShorten, err := s.repository.Create(shorten)

	if err != nil {
		return newShorten, nil
	}

	return newShorten, nil
}
