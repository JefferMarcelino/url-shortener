package domain

type ShortURL struct {
	Code    string
	LongURL string
}

type URLRepository interface {
	Save(url ShortURL) error
	GetByCode(code string) (*ShortURL, error)
}
