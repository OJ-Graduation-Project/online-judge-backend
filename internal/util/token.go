package util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SecretKey string = "yousseftarekaymanradwanverginia3agoog"

func CreateToken(issuer string) string {
	cookieExpiration := time.Now().Add(time.Hour * 3 * 24).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: cookieExpiration,
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		panic(err)
	}
	return token
}

func AuthenticateToken(token string) (string, error) { //returns issuer of token: email
	authToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		fmt.Println("USER IS UNAUTHENTICATED")
		return "", err
	}

	claims := authToken.Claims.(*jwt.StandardClaims)

	return claims.Issuer, nil
}
