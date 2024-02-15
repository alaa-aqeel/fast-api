package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(username string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": EXPIRE_TOKEN,
	})

	return accessToken.SignedString(SECRET_KEY)
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	var claims jwt.MapClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
