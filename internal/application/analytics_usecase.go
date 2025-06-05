package application

import "urlshortener/internal/domain"

type AnalyticsUseCase struct {
	analyticsReader domain.AnalyticsReader
}

func NewAnalyticsUseCase(analyticsReader domain.AnalyticsReader) *AnalyticsUseCase {
	return &AnalyticsUseCase{analyticsReader: analyticsReader}
}

func (h *AnalyticsUseCase) GetClickEventsByCode(code string) ([]domain.ClickEvent, error) {
	return h.analyticsReader.GetClickEventsByCode(code)
}
