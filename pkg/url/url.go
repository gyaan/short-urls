package url

import (
	"regexp"
)

type url struct{}

const RegexForValidUrl = `^((https?|ftp|smtp):\/\/)?(www.)?[a-z0-9]+\.[a-z]+(\/[a-zA-Z0-9#]+\/?)*$`

type Url interface {
	ValidateUrl(urlString string) bool
}

func New() Url {
	return &url{}
}

// ValidateUrl validate a given string is url or not
func (u url) ValidateUrl(urlString string) bool {
	validUrl := regexp.MustCompile(RegexForValidUrl)
	return validUrl.MatchString(urlString)
}
