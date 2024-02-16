package main

import (
	"github.com/alaa-aqeel/fast-api/api"
	"github.com/alaa-aqeel/fast-api/config"
	"github.com/alaa-aqeel/fast-api/pkg"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	router := router.New()
	api.ApiRouter(router.Group("/api/v1"))
	pkg.FastHttpServer(func(ctx *fasthttp.RequestCtx) {

		router.Handler(ctx)
	}, config.APP_URL)
}
