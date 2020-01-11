package handler

import (
	"fmt"
	"github.com/gyaan/short-urls/internal/repositories"
	"net/http"
)

type handler struct {
	shortUrlRepository repositories.ShortUrls
	userRepository     repositories.Users
}

type Handler interface {
	//common handler
	HomeHandler(w http.ResponseWriter, r *http.Request)

	//short url handler
	GetAShortUrl(w http.ResponseWriter, r *http.Request)
	GetAllShortUrl(w http.ResponseWriter, r *http.Request)
	DeleteShortUrl(w http.ResponseWriter, r *http.Request)
	CreateShortUrl(w http.ResponseWriter, r *http.Request)
	UpdateShortUrl(w http.ResponseWriter, r *http.Request)

	//user handler
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)

	//redirect short url to actual url
	RedirectToActualUrl(w http.ResponseWriter, r *http.Request)
}

//NewHandler creates new Handler
func NewHandler(shortUrlRepository repositories.ShortUrls, userRepository repositories.Users) Handler {
	return &handler{
		shortUrlRepository: shortUrlRepository,
		userRepository:     userRepository,
	}
}

//HomeHandler handles / urls (base url)
func (h *handler) HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Welcome to short-urls open source")

	if err != nil {
		fmt.Println("something went wrong!")
	}
}
