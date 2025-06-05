package usecase

import (
	"math/rand"
	"time"
	"urlshortener/internal/domain/model"
	"urlshortener/internal/domain/repository"
)

type URLUseCase struct {
	urlRepo       repository.URLRepository
	analyticsRepo repository.ClickAnalyticsRepository
}

func NewURLUseCase(urlRepo repository.URLRepository, analyticsRepo repository.ClickAnalyticsRepository) *URLUseCase {
	return &URLUseCase{urlRepo: urlRepo, analyticsRepo: analyticsRepo}
}

func (uc *URLUseCase) Shorten(url string) (string, error) {
	code := generateCode(6)

	newShortUrl := model.ShortURL{
		LongURL: url,
		Code:    code,
	}

	err := uc.urlRepo.Save(newShortUrl)

	return code, err
}

func (uc *URLUseCase) Resolve(code, ip, userAgent string) (string, error) {
	url, err := uc.urlRepo.GetByCode(code)
	if err != nil {
		return "", err
	}

	analytics := &model.ClickAnalytics{
		Code:      code,
		Timestamp: time.Now().Unix(),
		IP:        ip,
		UserAgent: userAgent,
	}

	go uc.analyticsRepo.Save(analytics)

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
