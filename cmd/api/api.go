package main

import (
	"regexp"

	"github.com/alaa-aqeel/fast-api/api"
	"github.com/alaa-aqeel/fast-api/config"
	"github.com/alaa-aqeel/fast-api/pkg"
	"github.com/alaa-aqeel/fast-api/pkg/utils"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	router := router.New()
	api.ApiRouter(router.Group("/api"))
	pkg.FastHttpServer(func(ctx *fasthttp.RequestCtx) {
		if isMatch, _ := regexp.Match("^/api/", ctx.Request.URI().PathOriginal()); isMatch {
			router.Handler(ctx)
			return
		}
		utils.Response(ctx).StatusCode(404).Send(utils.Map{
			"message": "not found !!",
		})
	}, config.APP_URL)
}
