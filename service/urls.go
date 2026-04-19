package service

import (
	"my-shortener/model"
	"my-shortener/repository"
)

type URLService interface {
	GetAllUrls() ([]*model.URL, error)

	Create(originalURL *model.URL) (*model.URL, error)
	GetNewUrlByShortCode(shortCode string) (*model.URL, error)
}

type urlService struct{}

func NewURLService() URLService {
	return &urlService{}
}

func (s *urlService) GetAllUrls() ([]*model.URL, error) {
	urlRepo := repository.NewUrlRepository()

	url, err := urlRepo.GetAllUrls()
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *urlService) Create(originalURL *model.URL) (*model.URL, error) {
	urlRepo := repository.NewUrlRepository()
	return urlRepo.Create(originalURL)
}

func (s *urlService) GetNewUrlByShortCode(shortCode string) (*model.URL, error) {
	urlRepo := repository.NewUrlRepository()
	return urlRepo.GetNewUrlByShortCode(shortCode)
}