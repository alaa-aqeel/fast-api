package main

import (
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/valyala/fasthttp"
)

func AuthUser(ctx *fasthttp.RequestCtx) jwt.MapClaims {
	accessToken := string(ctx.Request.Header.Peek("Authorization"))
	if accessToken == "" {
		return nil
	}
	tk := strings.Split(accessToken, " ")
	claims, err := VerifyAccessToken(tk[1])
	if err != nil {

		return nil
	}

	return claims
}
