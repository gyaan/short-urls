package url_shortner

import "strconv"

type urlShortener struct {
	charMap map[int]string
}

type UrlShortener interface {
	GetShortUrl(urlIdentifier int) string
}

func New() UrlShortener {
	return &urlShortener{
		charMap: getCharMap(),
	}
}

func (s *urlShortener) GetShortUrl(urlIdentifier int) string {
	return s.convertNumberToShortUrl(urlIdentifier)
}

func getCharMap() map[int]string {

	chrMap := map[int]string{}
	mapIndex := 0

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

func (s urlShortener) convertNumberToShortUrl(n int) string {
	srtUlr := ""
	for n > 0 {
		srtUlr += s.charMap[n%62]
		n = n / 62
	}
	return srtUlr
}

func (s urlShortener) convertShortUrlToNumber(str string) int {

	num := 0
	/*
		if ('a' <= shortURL[i] && shortURL[i] <= 'z')
		id = id*62 + shortURL[i] - 'a';
		if ('A' <= shortURL[i] && shortURL[i] <= 'Z')
		id = id*62 + shortURL[i] - 'A' + 26; a := 0
		if ('0' <= shortURL[i] && shortURL[i] <= '9')
		id = id*62 + shortURL[i] - '0' + 52;
	*/

	for i := 0; i < len(str); i++ {

		if 'a' <= str[i] && str[i] <= 'z' {
			num = num*62 + int(str[i]) - int('a')
		}

		if 'A' <= str[i] && str[i] <= 'Z' {
			num = num*62 + int(str[i]) - int('A') + 26
		}

		if '0' <= str[i] && str[i] <= '9' {
			num = num*62 + int(str[i]) - int('0') + 52
		}

	}
	return num
}
