package access_token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

//Claims
type Claims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

// GetToken returns a token for verified user
// use this function after user verification
func GetToken(UserId string, tokenExpiryTime int64, jwtSecret string) (string, error) {

	//set expiration time for token
	expirationTime := time.Now().Add(time.Duration(tokenExpiryTime) * time.Minute)

	//create claim
	claims := &Claims{
		Id: UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	//creat token with hs256 method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//get token string
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Printf("error generating access access_token")
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validate a access token and return claims
func ValidateToken(tokenString string, jwtSecret string) (map[string]interface{}, error) {

	//parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	//error parsing token
	if err != nil {
		return nil, err
	}

	//Get the claim and verify token is valid or not
	claims, ok := token.Claims.(jwt.MapClaims)

	//issue with claims or token is not valid
	if !ok || !token.Valid {
		return nil, errors.New("access token is not a valid token")
	}

	//everything is fine
	return claims, nil
}
