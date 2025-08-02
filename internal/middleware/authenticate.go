package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gyaan/short-urls/internal/access_token"
	"github.com/gyaan/short-urls/internal/config"
)

// Authenticate validates the access token passed in the request
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")

		// Check if authorization header is present
		if authToken == "" {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		// Validate Bearer token format
		parts := strings.Split(authToken, "Bearer")
		if len(parts) != 2 {
			http.Error(w, "Invalid bearer token format", http.StatusBadRequest)
			return
		}

		token := strings.TrimSpace(parts[1])
		claims, err := access_token.ValidateToken(token, config.GetConf().JWTSecret)
		if err != nil {
			log.Printf("Failed to validate access token: %v", err)
			http.Error(w, "Invalid access token", http.StatusUnauthorized)
			return
		}

		// Add user ID to context for use in handlers
		ctx := context.WithValue(r.Context(), "user_id", claims["id"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
