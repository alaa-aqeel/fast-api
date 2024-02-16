package utils

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
