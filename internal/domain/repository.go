package domain

type URLRepository interface {
	Save(url URL) error
	GetByCode(code string) (*URL, error)
}
