package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gyaan/short-urls/internal/access-token"
	"github.com/gyaan/short-urls/internal/models"
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
	Id          string `json:"id"`
	AccessToken string `json:"access_token"`
}

//GetAccessToken returns the access token after user verification
func (h *handler) GetAccessToken(w http.ResponseWriter, r *http.Request) {
	var accessTokenRequest AccessTokenRequest
	err := json.NewDecoder(r.Body).Decode(&accessTokenRequest)
	errResponse := models.ErrorResponse{ErrorMessage: "error in generating access token", Retry: false}

	if err != nil {
		log.Printf("error decoding get access access-token request %v", err)
		http.Error(w, errResponse.Error(), http.StatusBadRequest)
		return
	}

	if len(accessTokenRequest.Password) <= 0 || len(accessTokenRequest.Name) <= 0 {
		errResponse.ErrorMessage = "name and password are required to generate access token"
		http.Error(w, errResponse.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Generating access-token for %s", accessTokenRequest.Name)

	//get user details
	user, err := h.userRepository.GetUserDetailsByName(r.Context(), accessTokenRequest.Name)
	if user == nil {
		log.Printf("wrong credentials for access access-token %v", err)
		errResponse.ErrorMessage = "user not found."
		http.Error(w, errResponse.Error(), http.StatusBadRequest)
		return
	}

	//compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(accessTokenRequest.Password))
	if err != nil {
		log.Printf("wrong credentials for access access-token, %v", err)
		errResponse.ErrorMessage = "wrong username password."
		http.Error(w, errResponse.Error(), http.StatusBadRequest)
		return
	}

	tokenString, err := access_token.GetToken(user.ID.Hex())
	if err != nil {
		log.Printf("error in access access-token generation %v", err)
		http.Error(w, errResponse.Error(), http.StatusBadRequest)
		return
	}
	accessTokenRes := AccessTokenResponse{
		Id:          user.ID.Hex(),
		AccessToken: tokenString,
	}

	bytes, err := json.Marshal(accessTokenRes)
	if err != nil {
		log.Printf("error marshaling jwt access-token response %v", err)
		http.Error(w, errResponse.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bytes))
}
