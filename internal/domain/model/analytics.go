package model

type ClickAnalytics struct {
	Code      string
	Timestamp int64
	IP        string
	UserAgent string
}

type Analytics struct {
	IP        string
	UserAgent string
	Timestamp string
}
