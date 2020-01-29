package middleware

import (
	"context"
	"github.com/gyaan/short-urls/internal/access_token"
	"github.com/gyaan/short-urls/internal/config"
	"github.com/gyaan/short-urls/internal/models"
	"log"
	"net/http"
	"strings"
)

// Authenticate validate access token passed in request
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errResponse := models.ErrorResponse{ErrorMessage: "Error with access token.", Retry: false}
		reqToken := r.Header.Get("Authorization")

		//check if authorization code is there or not
		if len(reqToken) == 0 {
			errResponse.ErrorMessage = "unauthorized access"
			http.Error(w, errResponse.Error(), http.StatusBadRequest)
			return
		}

		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) != 2 {
			errResponse.ErrorMessage = "invalid bearer authorization access token."
			http.Error(w, errResponse.Error(), http.StatusBadRequest)
			return
		}
		reqToken = strings.TrimSpace(splitToken[1])
		claims, err := access_token.ValidateToken(reqToken,config.GetConf().JWTSecret)

		if err != nil {
			log.Println("error in access token validation")
			errResponse.ErrorMessage = "error in access token validation"
			http.Error(w, errResponse.Error(), http.StatusInternalServerError)
			return
		}

		//add user_id to context for later use in handler
		newContext := context.WithValue(r.Context(), "user_id", claims["id"])
		next.ServeHTTP(w, r.WithContext(newContext))
	})
}
