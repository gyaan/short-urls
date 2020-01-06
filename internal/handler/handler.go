package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gyaan/short-urls/internal/database"
	"github.com/gyaan/short-urls/internal/models"
	"net/http"
)

type Handler struct {
	database database.Database
}

func NewHandler(database2 database.Database) *Handler {
	return &Handler{
		database: database2,
	}
}
func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Welcome to short-urls open source")

	if err != nil {
		fmt.Println(" error sending response")
	}

}

func (h *Handler) GetAShortUrl(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Implement get a short url handler")

	if err != nil {
		fmt.Println(" error sending response")
	}

}

func (h *Handler) GetAllShortUrl(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	urls, err := h.database.GetAllShortUrls(context.Background())

	if err != nil {
		_, err = fmt.Fprintf(w, "unable to get shor urls")
	}

	bytes, err := json.Marshal(urls)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (h *Handler) DeleteShortUrl(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Implement delete short urls handler")

	if err != nil {
		fmt.Println(" error sending response")
	}

}

func (h *Handler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {

	err := h.database.CreateShortUrl(context.Background(), models.ShortUrl{
		Url: "http://google.com",
	})

	if err != nil {
		_, err = fmt.Fprintf(w, "unable to create shor urls")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "short url created!")
}

func (h *Handler) UpdateShortUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, vars["short_url_id"])
}
