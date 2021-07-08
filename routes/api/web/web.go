package web

import (
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API"
)

func LoadWeb(router *gin.Engine) {
	noAuthAPI := router.Group("/api/Web")
	{
		noAuthAPI.GET("/test", API.Test)
	}
}
