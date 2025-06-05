package usecase

import (
	"urlshortener/internal/domain/model"
	"urlshortener/internal/domain/repository"
)

type AnalyticsUseCase struct {
	analyticsRepo repository.ClickAnalyticsRepository
}

func NewAnalyticsUseCase(analyticsRepo repository.ClickAnalyticsRepository) *AnalyticsUseCase {
	return &AnalyticsUseCase{analyticsRepo: analyticsRepo}
}

func (h *AnalyticsUseCase) GetAnalyticsByCode(code string) ([]model.Analytics, error) {
	return h.analyticsRepo.GetAnalyticsByCode(code)
}
