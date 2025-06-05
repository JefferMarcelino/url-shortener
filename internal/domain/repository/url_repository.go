package repository

import "urlshortener/internal/domain/model"

type URLRepository interface {
	Save(url model.ShortURL) error
	GetByCode(code string) (*model.ShortURL, error)
}
