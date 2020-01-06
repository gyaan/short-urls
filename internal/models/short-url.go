package models

type ShortUrl struct {
	Url    string `json:"actual_url"`
	NewUrl string `json:"new_url"`
}
