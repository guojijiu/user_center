package web

import (
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Web"
	"user_center/app/Http/Middleware"
)

func LoadWeb(router *gin.Engine) {

	noAuthAPI := router.Group("/api/web", Middleware.Middleware.Api...)
	{
		// 获取验证码
		noAuthAPI.GET("/captcha", Web.UserController{}.GetCaptcha)
		// 发送注册激活码
		noAuthAPI.POST("/send_register_code", Web.UserController{}.SendRegisterCode)
	}
}
