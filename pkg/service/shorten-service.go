package service

import (
	"github.com/teris-io/shortid"
	"github.com/zekhoi/learn-golang/pkg/entity"
	"github.com/zekhoi/learn-golang/pkg/repository"
	"github.com/zekhoi/learn-golang/pkg/request"
)

type ShortenService interface {
	CreateShorten(request request.CreateShortenRequest) (entity.Shorten, error)
	GetShortenByCode(code string) (entity.Shorten, error)
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

func (s *shortenService) GetShortenByCode(code string) (entity.Shorten, error) {
	shorten, err := s.repository.FindByCode(code)

	if err != nil {
		return shorten, err
	}

	return shorten, nil
}

func (s *shortenService) CreateShorten(request request.CreateShortenRequest) (entity.Shorten, error) {

	shorten := entity.Shorten{}
	shorten.OriginalUrl = request.OriginalUrl
	shorten.ShortUrl = shortid.MustGenerate()
	newShorten, err := s.repository.Create(shorten)

	if err != nil {
		return newShorten, err
	}

	return newShorten, nil
}
