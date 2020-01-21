package url_shortner

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {

	tests := []struct {
		name string
		want UrlShortener
	}{
		{
			name: "valid",
			want: &urlShortener{
				charMap: getCharMap(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCharMap(t *testing.T) {
	tests := []struct {
		name string
		want map[int64]string
	}{
		{
			name: "valid",
			want: map[int64]string{
				0:  "a",
				1:  "b",
				2:  "c",
				3:  "d",
				4:  "e",
				5:  "f",
				6:  "g",
				7:  "h",
				8:  "i",
				9:  "j",
				10: "k",
				11: "l",
				12: "m",
				13: "n",
				14: "o",
				15: "p",
				16: "q",
				17: "r",
				18: "s",
				19: "t",
				20: "u",
				21: "v",
				22: "w",
				23: "x",
				24: "y",
				25: "z",
				26: "A",
				27: "B",
				28: "C",
				29: "D",
				30: "E",
				31: "F",
				32: "G",
				33: "H",
				34: "I",
				35: "J",
				36: "K",
				37: "L",
				38: "M",
				39: "N",
				40: "O",
				41: "P",
				42: "Q",
				43: "R",
				44: "S",
				45: "T",
				46: "U",
				47: "V",
				48: "W",
				49: "X",
				50: "Y",
				51: "Z",
				52: "0",
				53: "1",
				54: "2",
				55: "3",
				56: "4",
				57: "5",
				58: "6",
				59: "7",
				60: "8",
				61: "9",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCharMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCharMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Example_urlShortener_getCharMap() {
	fmt.Print(getCharMap())
	//Output: map[0:a 1:b 2:c 3:d 4:e 5:f 6:g 7:h 8:i 9:j 10:k 11:l 12:m 13:n 14:o 15:p 16:q 17:r 18:s 19:t 20:u 21:v 22:w 23:x 24:y 25:z 26:A 27:B 28:C 29:D 30:E 31:F 32:G 33:H 34:I 35:J 36:K 37:L 38:M 39:N 40:O 41:P 42:Q 43:R 44:S 45:T 46:U 47:V 48:W 49:X 50:Y 51:Z 52:0 53:1 54:2 55:3 56:4 57:5 58:6 59:7 60:8 61:9]
}

func Test_urlShortener_GetShortUrl(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		s    urlShortener
		args args
		want string
	}{
		{
			name: "valid",
			s: urlShortener{
				charMap: getCharMap(),
			},
			args: args{
				1212,
			},
			want: "tI",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetShortUrl(tt.args.n); got != tt.want {
				t.Errorf("urlShortener.GetShortUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Example_urlShortener_GetShortUrl() {
	fmt.Printf(urlShortener{charMap: getCharMap()}.GetShortUrl(1212))
	//Output: tI
}

func Test_urlShortener_GetIdentifierNumberFromShortUrl(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		s    urlShortener
		args args
		want int64
	}{
		{
			name: "valid",
			s: urlShortener{
				charMap: getCharMap(),
			},
			args: args{
				"tI",
			},
			want: int64(1212),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetIdentifierNumberFromShortUrl(tt.args.str); got != tt.want {
				t.Errorf("urlShortener.GetIdentifierNumberFromShortUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleUrlShortener_GetIdentifierNumberFromShortUrl() {
	fmt.Printf("%d", urlShortener{charMap: getCharMap()}.GetIdentifierNumberFromShortUrl("tI"))
	//Output: 1212
}
