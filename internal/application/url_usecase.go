package application

import (
	"math/rand"
	"urlshortener/internal/domain"
)

type URLUseCase struct {
	urlRepo           domain.URLRepository
	analyticsReporter domain.AnalyticsReporter
}

func NewURLUseCase(urlRepo domain.URLRepository, analyticsReporter domain.AnalyticsReporter) *URLUseCase {
	return &URLUseCase{urlRepo: urlRepo, analyticsReporter: analyticsReporter}
}

func (uc *URLUseCase) Shorten(url string) (string, error) {
	code := generateCode(6)

	newShortUrl := domain.ShortURL{
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

	clickEvent := &domain.ClickEvent{
		Code:      code,
		IP:        ip,
		UserAgent: userAgent,
	}

	go uc.analyticsReporter.Save(clickEvent)

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
