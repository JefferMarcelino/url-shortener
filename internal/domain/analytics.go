package domain

type ClickEvent struct {
	Code      string
	Timestamp string
	IP        string
	UserAgent string
}

type AnalyticsReporter interface {
	Save(clickEvent *ClickEvent) error
}

type AnalyticsReader interface {
	GetClickEventsByCode(code string) ([]ClickEvent, error)
}
