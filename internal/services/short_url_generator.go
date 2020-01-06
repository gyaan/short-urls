package services

type shortUrlService struct {
}

type ShortUrlService interface {
	GetShortUrl(str string) string
}

func NewShortUrlService() ShortUrlService {
	return &shortUrlService{}
}

func (s *shortUrlService) GetShortUrl(str string) string {
	return str
}
