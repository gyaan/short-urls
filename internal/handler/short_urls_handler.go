package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gyaan/short-urls/internal/models"
	"github.com/gyaan/short-urls/pkg/url"
	"log"
	"net/http"
)

//CreateShortUrlRequest
type CreateShortUrlRequest struct {
	Url string `json:"url"`
}

//UpdateShortUrlRequest
type UpdateShortUrlRequest struct {
	Status int `json:"status"`
}

//CreateShortUrl creates new short url
func (h *handler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var srtUrlReq CreateShortUrlRequest

	//get values
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&srtUrlReq)
	if err != nil {
		log.Println("Error with create short url request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validUrl := url.New().ValidateUrl(srtUrlReq.Url)
	if !validUrl {
		log.Printf("url validation failed for url %s", srtUrlReq.Url)
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}

	srtUrl, err := h.shortUrlRepository.CreateShortUrl(r.Context(), srtUrlReq.Url)
	if err != nil {
		log.Printf("Error while creating new short urls %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(srtUrl)
	if err != nil {
		log.Printf("Error marshaling short url")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bytes))
}

//GetAShortUrl returns a single url details using url id
func (h *handler) GetAShortUrl(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	log.Printf("Get short url details for %s", vars["short_url_id"])

	srtUrl, err := h.shortUrlRepository.GetAShortUrl(r.Context(), vars["short_url_id"])
	if err != nil {
		log.Printf("Error fetching short url details for %s", vars["short_url_id"])
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(srtUrl)
	if err != nil {
		log.Printf("Error json marshaling for short url %v", srtUrl)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bytes))
}

//GetAllShortUrl returns all urls
func (h *handler) GetAllShortUrl(w http.ResponseWriter, r *http.Request) {
	srtUrls, err := h.shortUrlRepository.GetAllShortUrls(r.Context())

	if err != nil {
		log.Printf("Error feching all short url details")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(srtUrls)
	if err != nil {
		log.Printf("Error marshaling all short url response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bytes))
}

//DeleteShortUrl deletes a url using url id
func (h *handler) DeleteShortUrl(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	err := h.shortUrlRepository.DeleteShortUrl(r.Context(), vars["short_url_id"])

	if err != nil {
		log.Printf("Error deleting shor url for id %s", vars["short_url_id"])
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "successfully deleted short url entry")
}

//UpdateShortUrl updates existing url using url id
func (h *handler) UpdateShortUrl(w http.ResponseWriter, r *http.Request) {
	var srtUrlUpdateRequest UpdateShortUrlRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&srtUrlUpdateRequest)

	if err != nil {
		log.Printf("Error with update short url request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	err = h.shortUrlRepository.UpdateShortUrls(r.Context(), vars["short_url_id"], models.ShortUrl{Status: int32(srtUrlUpdateRequest.Status)})

	if err != nil {
		log.Printf("Error updating short url details for short url id %v", vars["short_url_id"])
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "successfully updated url details")
}

//RedirectToActualUrl redirect short urls to actual url
func (h *handler) RedirectToActualUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("Redirecting to actual url for short url %s", vars["short_url"])

	shortUrl, err := h.shortUrlRepository.GetActualUrlOfAShortUrl(r.Context(), vars["short_url"])

	if err != nil {
		log.Printf("Error with url redirection for short url %s", vars["short_url"])
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, shortUrl.Url, http.StatusSeeOther)
}
