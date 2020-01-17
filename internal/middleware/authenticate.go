package middleware

import (
	"errors"
	"github.com/gyaan/short-urls/internal/access-token"
	"log"
	"net/http"
	"strings"
)

// Authenticate validate access token passed in request
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")

		//check if authorization code is there or not
		if len(reqToken) == 0 {
			http.Error(w, errors.New("unauthorized access").Error(), http.StatusBadRequest)
			return
		}

		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) != 2 {
			http.Error(w, errors.New("invalid bearer authorization access-token 1").Error(), http.StatusBadRequest)
			return
		}
		reqToken = strings.TrimSpace(splitToken[1])
		isValid, err := access_token.ValidateToken(reqToken)
		if !isValid || err != nil {
			log.Println("error in access-token validation")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
