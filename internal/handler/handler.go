package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gyaan/short-urls/internal/models"
	"github.com/gyaan/short-urls/internal/repositories"
	"net/http"
)

type handler struct {
	database repositories.ShortUrls
}

//Handler
type Handler interface {
	HomeHandler(w http.ResponseWriter, r *http.Request)
	GetAShortUrl(w http.ResponseWriter, r *http.Request)
	GetAllShortUrl(w http.ResponseWriter, r *http.Request)
	DeleteShortUrl(w http.ResponseWriter, r *http.Request)
	CreateShortUrl(w http.ResponseWriter, r *http.Request)
	UpdateShortUrl(w http.ResponseWriter, r *http.Request)
}

//NewHandler
func NewHandler(database2 repositories.ShortUrls) Handler {
	return &handler{
		database: database2,
	}
}

//HomeHandler
func (h *handler) HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Welcome to short-urls open source")

	if err != nil {
		fmt.Println(" error sending response")
	}

}

//GetAShortUrl
func (h *handler) GetAShortUrl(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Implement get a short url handler")

	if err != nil {
		fmt.Println(" error sending response")
	}

}

//GetAllShortUrl
func (h *handler) GetAllShortUrl(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	urls, err := h.database.GetAllShortUrls(context.Background())

	if err != nil {
		_, err = fmt.Fprintf(w, "unable to get shor urls")
	}

	bytes, err := json.Marshal(urls)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

//DeleteShortUrl
func (h *handler) DeleteShortUrl(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Implement delete short urls handler")

	if err != nil {
		fmt.Println(" error sending response")
	}

}

//CreateShortUrl
func (h *handler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {

	err := h.database.CreateShortUrl(context.Background(), models.ShortUrl{
		Url: "http://google.com",
	})

	if err != nil {
		_, err = fmt.Fprintf(w, "unable to create shor urls")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "short url created!")
}

//UpdateShortUrl
func (h *handler) UpdateShortUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, vars["short_url_id"])
}
