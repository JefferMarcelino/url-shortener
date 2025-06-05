package url

type ShortenRequest struct {
	OriginalURL string `json:"url"`
}

type ShortenResponse struct {
	ShortCode    string `json:"shortCode"`
	ShortURL     string `json:"shortUrl"`
	AnalyticsURL string `json:"analyticsUrl"`
	OriginalURL  string `json:"originalUrl"`
}
