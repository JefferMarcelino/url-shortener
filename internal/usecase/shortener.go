package usecase

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
	"urlshortener/internal/domain"
)

type ShortenerUseCase struct {
	repo domain.URLRepository
}

func NewShortenerUseCase(r domain.URLRepository) *ShortenerUseCase {
	return &ShortenerUseCase{repo: r}
}

func (s *ShortenerUseCase) Shorten(url string) (string, error) {
	code := generateShortCode()
	err := s.repo.Save(domain.URL{LongURL: url, Code: code})
	return code, err
}

func (s *ShortenerUseCase) Resolve(code string) (string, error) {
	url, err := s.repo.GetByCode(code)

	if err != nil {
		return "", err
	}

	return url.LongURL, nil
}

func generateShortCode() string {
	b := make([]byte, 4)
	rand.Read(b)
	return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")
}
