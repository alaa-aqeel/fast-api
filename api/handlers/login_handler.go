package handlers

import (
	"github.com/alaa-aqeel/fast-api/pkg/utils"
	"github.com/valyala/fasthttp"
)

var FakeUser utils.Map = utils.Map{
	"username": "admin",
	"password": "admin",
}

func LoginHandler(ctx *fasthttp.RequestCtx) {
	utils.Response(ctx).Send(utils.Map{
		"username": "invalid value",
	})
	// data := utils.ParseBody(ctx)
	// if data == nil {
	// 	utils.Response(ctx).Send(utils.Map{
	// 		"username": "must be not nil",
	// 	})
	// 	return
	// }
	// username := data.Get("username").(string)
	// if username == "" || username != FakeUser["username"] {
	// 	utils.Response(ctx).Send(utils.Map{
	// 		"username": "invalid value",
	// 	})
	// 	return
	// }
	// tk, _ := auth.Auth().CreateToken(username)
	// utils.Response(ctx).Send(utils.Map{
	// 	"access_token": tk,
	// 	"username":     username,
	// })
}
