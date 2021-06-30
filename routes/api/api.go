package api

import (
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API"
)

func LoadApi(router *gin.Engine) {
	noAuthAPI := router.Group("/api")
	{
		noAuthAPI.GET("/test", API.Test)
	}
}
