package models

import "time"

type URL struct {
	OriginalUrl  string
	ShortUrlPath string
	ExpiresAt    time.Time
}

type ShortenRequestModel struct {
	Url       string    `json:"url"`
	ExpiresAt time.Time `json:"expires_at"`
}

type ShortenResponseModel struct {
	ShortUrlPath string `json:"shorturlpath"`
}

type RedirectRequestModel struct {
	ShortUrlPath string `json:"shorturlpath"`
}

type RedirectResponseModel struct {
	Url string `json:"redirecturl"`
}
