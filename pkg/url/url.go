package url

type url struct{}

type Url interface {
	ValidateUrl(url string) error
}

func New() Url {
	return &url{}
}

func (u *url) ValidateUrl(url string) error {
	return nil
}
