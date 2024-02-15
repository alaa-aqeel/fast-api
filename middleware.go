package main

import "github.com/valyala/fasthttp"

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
