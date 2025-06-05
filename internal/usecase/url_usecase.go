package usecase

import (
	"math/rand"
	"urlshortener/internal/domain/model"
	"urlshortener/internal/domain/repository"
)

type URLUseCase struct {
	repo repository.URLRepository
}

func NewURLUseCase(r repository.URLRepository) *URLUseCase {
	return &URLUseCase{repo: r}
}

func (uc *URLUseCase) Shorten(url string) (string, error) {
	code := generateCode(6)

	newShortUrl := model.ShortURL{
		LongURL: url,
		Code:    code,
	}

	err := uc.repo.Save(newShortUrl)

	return code, err
}

func (uc *URLUseCase) Resolve(code string) (string, error) {
	url, err := uc.repo.GetByCode(code)

	if err != nil {
		return "", err
	}

	return url.LongURL, nil
}

func generateCode(length int) string {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result)
}
