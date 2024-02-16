package handlers

import (
	"github.com/alaa-aqeel/fast-api/pkg/utils"
	"github.com/alaa-aqeel/fast-api/pkg/utils/auth"
	"github.com/valyala/fasthttp"
)

func GetProfile(ctx *fasthttp.RequestCtx) {
	user := auth.Auth().GetClaimsFromToken(ctx)

	utils.Response(ctx).Send(utils.Map{
		"data": utils.Map{
			"username": user["sub"],
		},
	})
}
