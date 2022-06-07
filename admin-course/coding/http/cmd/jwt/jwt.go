package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type customClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var secret = []byte("this is a secure secret")

func create(username string) (string, error) {
	claims := customClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "golang-school",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)

	if err != nil {
		fmt.Println(err)
	}

	return signedToken, err
}

func validate(tokenString string) (jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["username"])
	} else {
		fmt.Println(err)
	}

	return *token, err
}
