package api

import (
	"github.com/alaa-aqeel/fast-api/api/handlers"
	"github.com/alaa-aqeel/fast-api/api/middlewares"
	"github.com/fasthttp/router"
)

func ApiRouter(r *router.Group) {
	r.POST("/login", handlers.LoginHandler)
	r.GET("/profile", middlewares.AuthMiddleware(handlers.GetProfile))
}
