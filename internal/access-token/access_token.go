package access_token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gyaan/short-urls/internal/config"
	"log"
	"time"
)

//Claims
type Claims struct {
	Name string `json:"username"`
	jwt.StandardClaims
}

// GetToken returns a token for verified user
// use this function after user verification
func GetToken(name string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(config.GetConf().TokenExpiryTime) * time.Minute)
	claims := &Claims{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetConf().JWTSecret))
	if err != nil {
		log.Printf("error generating access access-token")
		return "", err
	}
	return tokenString, nil
}

// ValidateToken validate a access token
func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetConf().JWTSecret), nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Printf("access-token validated for user %s", claims["name"])
		return true, nil
	} else {
		return false, err
	}
}
