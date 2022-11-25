package Middleware

import "github.com/gin-gonic/gin"

type middleware struct {
	Def     []gin.HandlerFunc
	Api     []gin.HandlerFunc
	JWTAuth []gin.HandlerFunc
	CSRF    []gin.HandlerFunc
	Cors    []gin.HandlerFunc
}

var Middleware middleware

func Init() {
	Middleware.Def = []gin.HandlerFunc{
		Logger(),
	}
	Middleware.Api = []gin.HandlerFunc{
		ApiTestMiddleware(),
	}

	Middleware.JWTAuth = []gin.HandlerFunc{
		JWTMiddleware(),
	}

	Middleware.Cors = []gin.HandlerFunc{
		Cors(),
	}
}
