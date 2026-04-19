package repository

import (
	"my-shortener/model"
	"my-shortener/model/query"
)

type UrlRepository interface {
	GetAllUrls() ([]*model.URL, error)

	Create(originalURL *model.URL) (*model.URL, error)
	GetNewUrlByShortCode(shortCode string) (*model.URL, error)
}

type urlRepository struct{}

func NewUrlRepository() UrlRepository {
	return &urlRepository{}
}

func (r *urlRepository) GetAllUrls() ([]*model.URL, error) {
	u := query.URL

	url, err := u.Find()
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (r *urlRepository) Create(originalURL *model.URL) (*model.URL, error) {
	u := query.URL
	err := u.Create(originalURL)
	if err != nil {
		return nil, err
	}
	return originalURL, nil
}

func (r *urlRepository) GetNewUrlByShortCode(shortCode string) (*model.URL, error) {
	u := query.URL

	url, err := u.Where(u.ShortCode.Eq(shortCode)).First()
	if err != nil {
		return nil, err
	}

	return url, nil
}