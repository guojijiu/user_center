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
		// 注册
		noAuthAPI.PUT("/register", Web.UserController{}.Register)
		// 发送忘记验证码
		noAuthAPI.POST("/send_forget_code", Web.UserController{}.SendForgetCode)
		// 忘记验证码获取详情
		noAuthAPI.GET("/detail_by_forget", Web.UserController{}.DetailByForget)
		// 重置密码
		noAuthAPI.POST("/reset_passwd", Web.UserController{}.ResetPasswd)
		// 登录
		noAuthAPI.POST("/login", Web.AuthController{}.Login)
	}
	AuthAPI := router.Group("/api/web", Middleware.Middleware.Api...)
	{
		// 登出
		AuthAPI.POST("/logout", Web.AuthController{}.Logout)

	}
}
