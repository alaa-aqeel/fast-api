package main

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func LoginHandler(ctx *fasthttp.RequestCtx) {
	data := ParseBody(ctx)
	if data == nil {
		Response(ctx).Send(Map{
			"username": "must be not nil",
		})
		return
	}
	username := (GetValue(data, "username")).(string)
	if len(username) == 0 {
		Response(ctx).Send(Map{
			"username": "must be not nil",
		})
		return
	}
	tk, _ := CreateAccessToken(username)
	Response(ctx).Send(Map{
		"access_token": tk,
	})
}

func GetProfile(ctx *fasthttp.RequestCtx) {
	user := AuthUser(ctx)

	Response(ctx).Send(Map{
		"data": Map{
			"username": user["sub"],
		},
	})
}

func main() {
	router := router.New()
	router.POST("/login", LoginHandler)
	router.GET("/profile", AuthMiddleware(GetProfile))
	FastHttpServer(HandlerReqeusts(router.Handler))
}
