package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/valyala/fasthttp"
)

var EXPIRE_TOKEN = time.Now().Add(time.Hour * 24).Unix()
var SECRET_KEY []byte = []byte("107684d9d2b8664121807f3d654")

func (auth *sAuth) GetClaimsFromToken(ctx *fasthttp.RequestCtx) jwt.MapClaims {
	accessToken := string(ctx.Request.Header.Peek("Authorization"))
	if accessToken == "" {
		return nil
	}
	tk := strings.Split(accessToken, " ")
	claims, err := auth.VerifyToken(tk[1])
	if err != nil {

		return nil
	}

	return claims
}

func (auth *sAuth) CreateToken(username string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": EXPIRE_TOKEN,
	})

	return accessToken.SignedString(SECRET_KEY)
}

func (auth *sAuth) VerifyToken(tokenString string) (jwt.MapClaims, error) {
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
