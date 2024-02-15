package main

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	router := router.New()
	router.GET("/", func(ctx *fasthttp.RequestCtx) {
		Response(ctx).Send(Map{
			"message": "Hello world",
		})
	})
	FastHttpServer(HandlerReqeusts(router.Handler))
}
