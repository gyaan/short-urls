// Package url_shortner provides URL shortening functionality
package url_shortner

import (
	"strconv"
)

// urlShortener implements URL shortening logic
type urlShortener struct {
	charMap map[int64]string
}

// UrlShortener defines the interface for URL shortening operations
type UrlShortener interface {
	GetShortUrl(urlIdentifier int64) string
	GetIdentifierNumberFromShortUrl(str string) int64
}

// New creates a new URL shortener instance
func New() UrlShortener {
	return &urlShortener{
		charMap: getCharMap(),
	}
}

// getCharMap creates a character mapping for URL shortening
func getCharMap() map[int64]string {
	charMap := map[int64]string{}
	mapIndex := int64(0)

	// Map lowercase letters a-z
	for j := rune('a'); j <= rune('z'); j++ {
		charMap[mapIndex] = string(j)
		mapIndex++
	}

	// Map uppercase letters A-Z
	for k := rune('A'); k <= rune('Z'); k++ {
		charMap[mapIndex] = string(k)
		mapIndex++
	}

	// Map digits 0-9
	for i := 0; i <= 9; i++ {
		charMap[mapIndex] = strconv.Itoa(i)
		mapIndex++
	}

	return charMap
}

// GetShortUrl converts a numeric identifier to a short URL string
func (s urlShortener) GetShortUrl(n int64) string {
	shortUrl := ""
	for n > 0 {
		shortUrl += s.charMap[n%62]
		n = n / 62
	}
	return reverseString(shortUrl)
}

func (s urlShortener) GetIdentifierNumberFromShortUrl(str string) int64 {

	num := int64(0)
	for i := 0; i < len(str); i++ {

		if 'a' <= str[i] && str[i] <= 'z' {
			num = num*62 + int64(rune(str[i])) - int64(rune('a'))
		}

		if 'A' <= str[i] && str[i] <= 'Z' {
			num = num*62 + int64(rune(str[i])) - int64(rune('A')) + 26
		}

		if '0' <= str[i] && str[i] <= '9' {
			num = num*62 + int64(rune(str[i])) - int64(rune('0')) + 52
		}

	}
	return num
}

// reverseString reverses a string using Unicode-aware operations
func reverseString(str string) string {
	// Get Unicode code points
	n := 0
	runes := make([]rune, len(str))
	for _, r := range str {
		runes[n] = r
		n++
	}
	runes = runes[0:n]

	// Reverse the runes
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	// Convert back to UTF-8
	return string(runes)
}
