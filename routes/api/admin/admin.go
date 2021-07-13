package admin

import (
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Admin/UserController"
	"user_center/app/Http/Middleware"
)

func LoadAdmin(router *gin.Engine) {
	AuthAPI := router.Group("/api/admin", Middleware.Middleware.Api...)
	{
		user := AuthAPI.Group("/user")
		{
			// 保存用户数据
			user.POST("store", UserController.Store)
		}

	}
}
