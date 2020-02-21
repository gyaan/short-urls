package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/gyaan/short-urls/internal/models"
	"github.com/gyaan/short-urls/internal/repositories"
	"github.com/gyaan/short-urls/pkg/pagination"
	"github.com/gyaan/short-urls/pkg/url"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//CreateShortUrlRequest
type CreateShortUrlRequest struct {
	Url string `json:"url"`
}

//UpdateShortUrlRequest
type UpdateShortUrlRequest struct {
	Status int `json:"status"`
}

type shortUrlHandler struct {
	shortUrlRepository repositories.ShortUrls
}

type ShortUrlHandler interface {
	CreateShortUrl(w http.ResponseWriter, r *http.Request)
	GetAShortUrl(w http.ResponseWriter, r *http.Request)
	GetAllShortUrl(w http.ResponseWriter, r *http.Request)
	DeleteShortUrl(w http.ResponseWriter, r *http.Request)
	UpdateShortUrl(w http.ResponseWriter, r *http.Request)
	RedirectToActualUrl(w http.ResponseWriter, r *http.Request)
	HomeHandler(w http.ResponseWriter, r *http.Request)
}

func NewShortUrlHandler(urls repositories.ShortUrls) ShortUrlHandler {
	return  &shortUrlHandler{
		shortUrlRepository:urls,
	}
}
//CreateShortUrl creates new short url
func (h *shortUrlHandler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
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
func (h *shortUrlHandler) GetAShortUrl(w http.ResponseWriter, r *http.Request) {
	shortUrlId := chi.URLParam(r, "short_url_id")
	log.Printf("Get short url details for %s", shortUrlId)
	errResponse := models.ErrorResponse{ErrorMessage: "error fetching short url details", Retry: false}

	srtUrl, err := h.shortUrlRepository.GetAShortUrl(r.Context(), shortUrlId)
	if err != nil {
		log.Printf("Error fetching short url details for %s, %v", shortUrlId, err)
		http.Error(w, errResponse.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(srtUrl)
	if err != nil {
		log.Printf("Error json marshaling for short url %v,%v", srtUrl, err)
		http.Error(w, errResponse.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bytes))
}

//GetAllShortUrl returns all urls
func (h *shortUrlHandler) GetAllShortUrl(w http.ResponseWriter, r *http.Request) {
	errResponse := models.ErrorResponse{ErrorMessage: "error fetching all short url details", Retry: false}
	var srtUrls []models.ShortUrl
	page := 1
	limit := 10
	pageString := r.URL.Query().Get("page")

	if len(pageString) > 0 {
		p, err := strconv.Atoi(pageString)
		if err != nil {
			log.Printf("Error parsing page url parameter %v", err)
			http.Error(w, errResponse.Error(), http.StatusBadRequest)
			return
		}
		page = p
	}
	limitString := r.URL.Query().Get("limit")

	if len(limitString) > 0 {
		l, err := strconv.Atoi(limitString)
		if err != nil {
			log.Printf("Error parsing limit url parameter %v", err)
			http.Error(w, errResponse.Error(), http.StatusBadRequest)
			return
		}
		limit = l
	}

	//get count of short urls for requested user
	count, err := h.shortUrlRepository.GetTotalShortUrlsCount(r.Context())
	if err != nil {
		log.Printf("Error parsing limit url parameter %v", err)
		http.Error(w, errResponse.Error(), http.StatusInternalServerError)
		return
	}

	//query only if there is short urls
	if count > 0 {
		//offset to fetch the elements
		offset := limit * (page - 1)
		su, err := h.shortUrlRepository.GetAllShortUrls(r.Context(), offset, limit)

		if err != nil {
			log.Printf("Error feching all short url details %v", err)
			http.Error(w, errResponse.Error(), http.StatusInternalServerError)
			return
		}
		srtUrls = su
	}
	paginationObj := pagination.New(count, int64(page), srtUrls, limit)
	response, err := paginationObj.GetPagination()
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling all short url response %v", err)
		http.Error(w, errResponse.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bytes))
}

//DeleteShortUrl deletes a url using url id
func (h *shortUrlHandler) DeleteShortUrl(w http.ResponseWriter, r *http.Request) {

	shortUrlId := chi.URLParam(r, "short_url_id")
	err := h.shortUrlRepository.DeleteShortUrl(r.Context(), shortUrlId)
	errResponse := models.ErrorResponse{ErrorMessage: "error removing short urls", Retry: false}

	if err != nil {
		log.Printf("Error deleting shor url for id %s, %v", shortUrlId, err)
		http.Error(w, errResponse.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "successfully deleted short url entry")
}

//UpdateShortUrl updates existing url using url id
func (h *shortUrlHandler) UpdateShortUrl(w http.ResponseWriter, r *http.Request) {
	var srtUrlUpdateRequest UpdateShortUrlRequest

	errResponse := models.ErrorResponse{ErrorMessage: "error updating short url", Retry: false}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&srtUrlUpdateRequest)

	if err != nil {
		log.Printf("Error with update short url request %v", err)
		http.Error(w, errResponse.Error(), http.StatusBadRequest)
		return
	}

	shortUrlId := chi.URLParam(r, "short_url_id")
	err = h.shortUrlRepository.UpdateShortUrls(r.Context(), shortUrlId, models.ShortUrl{Status: int32(srtUrlUpdateRequest.Status)})

	if err != nil {
		log.Printf("Error updating short url details for short url id %v,%v", shortUrlId, err)
		http.Error(w, errResponse.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "successfully updated url details")
}

//RedirectToActualUrl redirect short urls to actual url
func (h *shortUrlHandler) RedirectToActualUrl(w http.ResponseWriter, r *http.Request) {
	srtUrl := chi.URLParam(r, "short_url")
	log.Printf("Redirecting to actual url for short url %s", srtUrl)

	shortUrl, err := h.shortUrlRepository.GetActualUrlOfAShortUrl(r.Context(), srtUrl)
	if err != nil {
		log.Printf("Error with url redirection for short url %s", srtUrl)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	redirectUrl := shortUrl.Url
	if !strings.Contains(shortUrl.Url, "http") { //todo move this to url validation
		redirectUrl = "http://" + redirectUrl
	}
	//update click count
	err = h.shortUrlRepository.IncrementClickCountOfShortUrl(r.Context(), srtUrl)
	if err != nil { //just log the error
		fmt.Printf("Error increasing clicks count %v", err)
	}

	http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
}

//HomeHandler handles / urls (base url)
func (h *shortUrlHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to short-urls open source")
}