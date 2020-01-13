package url_shortner

import (
	"strconv"
)

type urlShortener struct {
	charMap map[int64]string
}

//UrlShortener
type UrlShortener interface {
	GetShortUrl(urlIdentifier int64) string
	GetIdentifierNumberFromShortUrl(str string) int64
}

//New
func New() UrlShortener {
	return &urlShortener{
		charMap: getCharMap(),
	}
}

func getCharMap() map[int64]string {

	chrMap := map[int64]string{}
	mapIndex := int64(0)

	//a - z
	for j := rune('a'); j <= rune('z'); j++ {
		chrMap[mapIndex] = string(j)
		mapIndex++
	}

	// A -Z
	for k := rune('A'); k <= rune('Z'); k++ {
		chrMap[mapIndex] = string(k)
		mapIndex++
	}

	//0 - 9
	for i := 0; i <= 9; i++ {
		chrMap[mapIndex] = strconv.Itoa(i)
		mapIndex++
	}

	return chrMap
}

func (s urlShortener) GetShortUrl(n int64) string {
	srtUlr := ""
	for n > 0 {
		srtUlr += s.charMap[n%62]
		n = n / 62
	}
	return reversString(srtUlr)
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

func reversString(str string) string {

	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(str))
	for _, r := range str {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}
