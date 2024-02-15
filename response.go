package main

import (
	"github.com/valyala/fasthttp"
)

type sResponse struct {
	statusCode  int
	ctx         *fasthttp.RequestCtx
	contentType string
}

func Response(ctx *fasthttp.RequestCtx) *sResponse {
	return &sResponse{
		ctx:         ctx,
		statusCode:  fasthttp.StatusOK,
		contentType: "text/plain; charset=utf-8",
	}
}

func (resp *sResponse) StatusCode(code int) *sResponse {
	resp.statusCode = code
	return resp
}

func (resp *sResponse) SetHeader(key string, val any) *sResponse {
	resp.ctx.Response.Header.Set("Content-Type", resp.contentType)
	return resp
}

func (resp *sResponse) Send(data Map) {
	resp.ctx.Response.SetStatusCode(resp.statusCode)
	resp.ctx.Response.Header.Set("Content-Type", "application/json")
	resp.ctx.SetBody(ToJson(data))
}
