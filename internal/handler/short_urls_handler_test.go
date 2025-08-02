package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gyaan/short-urls/internal/models"
	"github.com/gyaan/short-urls/mocks"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_shortUrlHandler_CreateShortUrl(t *testing.T) {
	//request data
	requestData, err := json.Marshal(CreateShortUrlRequest{
		Url: "https://www.google.com",
	})

	//user details
	shortUrl := models.ShortUrl{
		Url:    "https://www.google.com/",
		NewUrl: "https://newshorturl.com/",
	}

	if err != nil {
		t.Fatal(err)
	}

	//create request
	req, err := http.NewRequest("POST", "/short-urls", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}

	//creat new response writer
	rr := httptest.NewRecorder()

	//user repository mock
	shortUrlRepository := mocks.ShortUrls{}
	shortUrlRepository.On("CreateShortUrl", context.TODO(), mock.Anything).Return(&shortUrl, nil)
	newShortUrlHandler := NewShortUrlHandler(&shortUrlRepository)
	handler := http.HandlerFunc(newShortUrlHandler.CreateShortUrl)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func Test_shortUrlHandler_GetAShortUrl(t *testing.T) {

	//create request
	req, err := http.NewRequest("GET", "/short-urls", bytes.NewBufferString(url.Values{"short_url_id": {`xyx`}}.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	//creat new response writer
	rr := httptest.NewRecorder()

	//user details
	shortUrl := models.ShortUrl{
		Url:    "https://www.google.com/",
		NewUrl: "https://newshorturl.com/",
	}

	//user repository mock
	shortUrlRepository := mocks.ShortUrls{}
	shortUrlRepository.On("GetAShortUrl", context.TODO(), mock.Anything).Return(&shortUrl, nil)

	newShortUrlHandler := NewShortUrlHandler(&shortUrlRepository)
	handler := http.HandlerFunc(newShortUrlHandler.GetAShortUrl)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func Test_shortUrlHandler_GetAllShortUrl(t *testing.T) {

	//create request
	req, err := http.NewRequest("GET", "/short-urls", bytes.NewBufferString(url.Values{"page": {`1`}, "limit": {`10`}}.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	//creat new response writer
	rr := httptest.NewRecorder()

	//user details
	shortUrl := []models.ShortUrl{{
		Url:    "https://www.google.com/",
		NewUrl: "https://newshorturl.com/",
	}}

	//user repository mock
	shortUrlRepository := mocks.ShortUrls{}
	shortUrlRepository.On("GetTotalShortUrlsCount", context.TODO()).Return(int64(10), nil)
	shortUrlRepository.On("GetAllShortUrls", context.TODO(), mock.Anything, mock.Anything).Return(shortUrl, nil)
	newShortUrlHandler := NewShortUrlHandler(&shortUrlRepository)
	handler := http.HandlerFunc(newShortUrlHandler.GetAllShortUrl)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func Test_shortUrlHandler_DeleteShortUrl(t *testing.T) {
	//create request
	req, err := http.NewRequest("GET", "/short-urls", bytes.NewBufferString(url.Values{"short_url_id": {`xyx`}}.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	//creat new response writer
	rr := httptest.NewRecorder()

	//user repository mock
	shortUrlRepository := mocks.ShortUrls{}
	shortUrlRepository.On("DeleteShortUrl", context.TODO(), mock.Anything).Return(nil)

	newShortUrlHandler := NewShortUrlHandler(&shortUrlRepository)
	handler := http.HandlerFunc(newShortUrlHandler.DeleteShortUrl)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func Test_shortUrlHandler_UpdateShortUrl(t *testing.T) {
	//request data
	requestData, err := json.Marshal(UpdateShortUrlRequest{
		Status: 1,
	})

	if err != nil {
		t.Fatal(err)
	}

	//create request
	req, err := http.NewRequest("POST", "/short-urls?short_url_id=xyx", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}

	//creat new response writer
	rr := httptest.NewRecorder()

	//user repository mock
	shortUrlRepository := mocks.ShortUrls{}
	shortUrlRepository.On("UpdateShortUrls", context.TODO(), mock.Anything, mock.Anything).Return(nil)
	newShortUrlHandler := NewShortUrlHandler(&shortUrlRepository)
	handler := http.HandlerFunc(newShortUrlHandler.UpdateShortUrl)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func Test_shortUrlHandler_RedirectToActualUrl(t *testing.T) {
	//create request
	req, err := http.NewRequest("GET", "/short-urls", bytes.NewBufferString(url.Values{"short_url": {`xyx`}}.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	//creat new response writer
	rr := httptest.NewRecorder()

	//user details
	shortUrl := models.ShortUrl{
		Url:    "https://www.google.com/",
		NewUrl: "https://newshorturl.com/",
	}

	//user repository mock
	shortUrlRepository := mocks.ShortUrls{}
	shortUrlRepository.On("GetActualUrlOfAShortUrl", context.TODO(), mock.Anything).Return(&shortUrl, nil)
	shortUrlRepository.On("IncrementClickCountOfShortUrl", context.TODO(), mock.Anything).Return(nil)

	newShortUrlHandler := NewShortUrlHandler(&shortUrlRepository)
	handler := http.HandlerFunc(newShortUrlHandler.RedirectToActualUrl)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}
}

func Test_shortUrlHandler_HomeHandler(t *testing.T) {
	//create request
	req, err := http.NewRequest("GET", "/short-urls", nil)
	if err != nil {
		t.Fatal(err)
	}

	//creat new response writer
	rr := httptest.NewRecorder()

	//user repository mock
	shortUrlRepository := mocks.ShortUrls{}
	newShortUrlHandler := NewShortUrlHandler(&shortUrlRepository)
	handler := http.HandlerFunc(newShortUrlHandler.HomeHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
