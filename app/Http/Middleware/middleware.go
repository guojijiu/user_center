package Middleware

import "github.com/gin-gonic/gin"

type middleware struct {
	Def     []gin.HandlerFunc
	Api     []gin.HandlerFunc
	BackJWT []gin.HandlerFunc
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

	Middleware.BackJWT = []gin.HandlerFunc{
		BackJWTMiddleware(),
	}

	Middleware.Cors = []gin.HandlerFunc{
		Cors(),
	}
}
