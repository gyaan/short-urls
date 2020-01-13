package url

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Url
	}{
		{
			name: "valid",
			want: &url{},
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

func Test_url_ValidateUrl(t *testing.T) {

	u := url{}
	type args struct {
		url string
	}
	tests := []struct {
		name string
		u    Url
		args args
		want bool
	}{
		{
			name: "valid url",
			u:    u,
			args: args{
				url: "http://www.google.com",
			},
			want: true,
		},
		{
			name: "invalid url",
			u:    u,
			args: args{
				url: "asdf",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.ValidateUrl(tt.args.url); got != tt.want {
				t.Errorf("url.ValidateUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Example_url_ValidateUrl() {
	u := url{}
	fmt.Print(u.ValidateUrl("http://www.google.com"))
	// Output: true
}
