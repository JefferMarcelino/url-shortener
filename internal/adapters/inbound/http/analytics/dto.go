package analytics

type ClickEventDTO struct {
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
	UserAgent string `json:"userAgent"`
}

type AnalyticsResponse struct {
	ShortCode   string          `json:"shortCode"`
	TotalClicks int             `json:"totalClicks"`
	Clicks      []ClickEventDTO `json:"clicks"`
}
