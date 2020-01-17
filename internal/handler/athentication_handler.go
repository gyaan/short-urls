package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gyaan/short-urls/internal/access-token"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

//AccessTokenRequest
type AccessTokenRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

//AccessTokenResponse
type AccessTokenResponse struct {
	Name        string `json:"name"`
	AccessToken string `json:"access_token"`
}

//GetAccessToken returns the access token after user verification
func (h *handler) GetAccessToken(w http.ResponseWriter, r *http.Request) {
	//verify user credentials and release access access-token

	var accessTokenRequest AccessTokenRequest
	err := json.NewDecoder(r.Body).Decode(&accessTokenRequest)

	if err != nil {
		log.Println("Error decoding get access access-token request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(accessTokenRequest.Password) <= 0 || len(accessTokenRequest.Name) <= 0 {
		log.Printf("access access-token request without name and password")
		http.Error(w, errors.New("access access-token request without name and password").Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Generating access-token for %s", accessTokenRequest.Name)

	//get user details
	user, err := h.userRepository.GetUserDetails(r.Context(), accessTokenRequest.Name)
	if user == nil {
		log.Printf("wrong credentials for access access-token")
		http.Error(w, errors.New("wrong credentials for access access-token").Error(), http.StatusBadRequest)
		return
	}

	//compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(accessTokenRequest.Password))
	if err != nil {
		log.Printf("wrong credentials for access access-token")
		http.Error(w, errors.New("wrong credentials for access access-token").Error(), http.StatusBadRequest)
		return
	}

	tokenString, err := access_token.GetToken(accessTokenRequest.Name)
	if err != nil {
		log.Printf("error in access access-token generation")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	accessTokenRes := AccessTokenResponse{
		Name:        user.Name,
		AccessToken: tokenString,
	}

	bytes, err := json.Marshal(accessTokenRes)
	if err != nil {
		log.Printf("error marshaling jwt access-token response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bytes))
}
