package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/gyaan/short-urls/internal/models"
	"github.com/gyaan/short-urls/internal/repositories"
	"github.com/gyaan/short-urls/pkg/pagination"
	"github.com/gyaan/short-urls/pkg/url"
)

// CreateShortUrlRequest represents the request body for creating a short URL
type CreateShortUrlRequest struct {
	Url string `json:"url"`
}

// UpdateShortUrlRequest represents the request body for updating a short URL
type UpdateShortUrlRequest struct {
	Status int `json:"status"`
}

// shortUrlHandler handles HTTP requests for short URL operations
type shortUrlHandler struct {
	shortUrlRepository repositories.ShortUrls
}

// ShortUrlHandler defines the interface for short URL HTTP handlers
type ShortUrlHandler interface {
	CreateShortUrl(w http.ResponseWriter, r *http.Request)
	GetAShortUrl(w http.ResponseWriter, r *http.Request)
	GetAllShortUrl(w http.ResponseWriter, r *http.Request)
	DeleteShortUrl(w http.ResponseWriter, r *http.Request)
	UpdateShortUrl(w http.ResponseWriter, r *http.Request)
	RedirectToActualUrl(w http.ResponseWriter, r *http.Request)
	HomeHandler(w http.ResponseWriter, r *http.Request)
}

// NewShortUrlHandler creates a new short URL handler instance
func NewShortUrlHandler(urls repositories.ShortUrls) ShortUrlHandler {
	return &shortUrlHandler{
		shortUrlRepository: urls,
	}
}

// CreateShortUrl handles the creation of new short URLs
func (h *shortUrlHandler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var shortUrlReq CreateShortUrlRequest

	// Decode request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&shortUrlReq); err != nil {
		log.Printf("Failed to decode create short URL request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate URL
	if !url.New().ValidateUrl(shortUrlReq.Url) {
		log.Printf("URL validation failed for: %s", shortUrlReq.Url)
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// Create short URL
	shortUrl, err := h.shortUrlRepository.CreateShortUrl(r.Context(), shortUrlReq.Url)
	if err != nil {
		log.Printf("Failed to create short URL: %v", err)
		http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		return
	}

	// Marshal response
	responseBytes, err := json.Marshal(shortUrl)
	if err != nil {
		log.Printf("Failed to marshal short URL response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(responseBytes))
}

// GetAShortUrl retrieves a single short URL by its ID
func (h *shortUrlHandler) GetAShortUrl(w http.ResponseWriter, r *http.Request) {
	shortUrlID := chi.URLParam(r, "short_url_id")
	log.Printf("Retrieving short URL details for ID: %s", shortUrlID)

	shortUrl, err := h.shortUrlRepository.GetAShortUrl(r.Context(), shortUrlID)
	if err != nil {
		log.Printf("Failed to fetch short URL details for ID %s: %v", shortUrlID, err)
		http.Error(w, "Failed to fetch short URL details", http.StatusInternalServerError)
		return
	}

	responseBytes, err := json.Marshal(shortUrl)
	if err != nil {
		log.Printf("Failed to marshal short URL response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(responseBytes))
}

// GetAllShortUrl retrieves all short URLs with pagination
func (h *shortUrlHandler) GetAllShortUrl(w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters
	page := 1
	limit := 10

	if pageString := r.URL.Query().Get("page"); pageString != "" {
		if p, err := strconv.Atoi(pageString); err != nil {
			log.Printf("Failed to parse page parameter: %v", err)
			http.Error(w, "Invalid page parameter", http.StatusBadRequest)
			return
		} else {
			page = p
		}
	}

	if limitString := r.URL.Query().Get("limit"); limitString != "" {
		if l, err := strconv.Atoi(limitString); err != nil {
			log.Printf("Failed to parse limit parameter: %v", err)
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		} else {
			limit = l
		}
	}

	// Get total count of short URLs for the user
	totalCount, err := h.shortUrlRepository.GetTotalShortUrlsCount(r.Context())
	if err != nil {
		log.Printf("Failed to get total short URLs count: %v", err)
		http.Error(w, "Failed to retrieve short URLs", http.StatusInternalServerError)
		return
	}

	var shortUrls []models.ShortUrl

	// Fetch short URLs only if there are any
	if totalCount > 0 {
		offset := limit * (page - 1)
		shortUrls, err = h.shortUrlRepository.GetAllShortUrls(r.Context(), offset, limit)
		if err != nil {
			log.Printf("Failed to fetch short URLs: %v", err)
			http.Error(w, "Failed to retrieve short URLs", http.StatusInternalServerError)
			return
		}
	}

	// Create pagination response
	paginationObj := pagination.New(totalCount, int64(page), shortUrls, limit)
	response, err := paginationObj.GetPagination()
	if err != nil {
		log.Printf("Failed to create pagination response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("Failed to marshal pagination response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(responseBytes))
}

// DeleteShortUrl deletes a url using url id
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

// UpdateShortUrl updates existing url using url id
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

// RedirectToActualUrl redirect short urls to actual url
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

// HomeHandler handles / urls (base url)
func (h *shortUrlHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to short-urls open source")
}
