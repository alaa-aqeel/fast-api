package main

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	router := router.New()
	router.GET("", func(ctx *fasthttp.RequestCtx) {

		Response(ctx).Text("Hello world").Send()
	})
	FastHttpServer(HandlerReqeusts(router.Handler))
}
