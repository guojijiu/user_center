package admin

import (
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Admin/UserController"
	"user_center/app/Http/Middleware"
)

func LoadAdmin(router *gin.Engine) {
	AuthAPI := router.Group("/api/admin", Middleware.Middleware.Api...)
	{
		// 用户相关
		user := AuthAPI.Group("/user")
		{
			// 保存用户数据
			user.POST("store", UserController.Store)
			// 更新用户数据
			user.PUT("update", UserController.Update)
			// 获取用户详情
			user.GET("detail", UserController.Detail)
			// 获取用户列表
			user.GET("list", UserController.GetList)
			// 禁用用户
			user.DELETE("forbidden", UserController.Forbidden)
		}
	}
}
