package main

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func ParseBody(ctx *fasthttp.RequestCtx) Map {
	var data Map
	if err := json.Unmarshal(ctx.PostBody(), &data); err != nil {
		return nil
	}
	return data
}

func GetValue(data Map, key string) any {
	val, ok := data[key]
	if ok {
		return val
	}
	return nil
}
