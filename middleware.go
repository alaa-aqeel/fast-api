package main

import "github.com/valyala/fasthttp"

func AuthMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		accessToken := string(ctx.Request.Header.Peek("Authorization"))
		if accessToken == "" {
			Response(ctx).StatusCode(fasthttp.StatusUnauthorized).Send(Map{
				"message": "Unauthorized",
			})
			return
		}
		_, err := VerifyAccessToken(accessToken[len("Bearer "):])
		if err != nil {
			Response(ctx).StatusCode(fasthttp.StatusUnauthorized).Send(Map{
				"message": "Unauthorized",
			})
			return
		}
		next(ctx)
	}
}

func HandlerReqeusts(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set(
			"Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE",
		)
		ctx.Response.Header.Set(
			"Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		)
		next(ctx)
	}
}
