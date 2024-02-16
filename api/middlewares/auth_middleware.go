package middlewares

import (
	"github.com/alaa-aqeel/fast-api/pkg/utils"
	"github.com/alaa-aqeel/fast-api/pkg/utils/auth"
	"github.com/valyala/fasthttp"
)

func AuthMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		claims := auth.Auth().GetClaimsFromToken(ctx)
		if claims == nil {
			utils.Response(ctx).StatusCode(fasthttp.StatusUnauthorized).Send(utils.Map{
				"message": "Unauthorized",
			})
			return
		}
		next(ctx)
	}
}
