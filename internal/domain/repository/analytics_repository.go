package repository

import "urlshortener/internal/domain/model"

type ClickAnalyticsRepository interface {
	Save(analytics *model.ClickAnalytics) error
	GetAnalyticsByCode(code string) ([]model.Analytics, error)
}
