package routes

import (
	"github.com/gin-gonic/gin"
	"io"
	"user_center/app"
	"user_center/app/Http/Middleware"
	"user_center/config"
	"user_center/pkg/glog"
	"user_center/routes/api"
	"user_center/routes/api/admin"
	"user_center/routes/api/web"
)

func InitRouter() *gin.Engine {
	var router *gin.Engine
	gin.SetMode(config.GinModel)
	if app.InTest || app.InConsole {
		gin.DisableConsoleColor()
		gin.DefaultWriter = io.MultiWriter(glog.Channel("gin").Out)
	} else {
		// 非测试或命令将输出路由信息到屏幕上
		//gin.ForceConsoleColor()
		//gin.DefaultWriter = io.MultiWriter(glog.Channel("gin").Out, os.Stdout)
	}

	router = gin.New()
	router.Use(gin.RecoveryWithWriter(glog.Channel("gin").Out), gin.Logger())
	// 加载默认中间件
	router.Use(Middleware.Middleware.Def...)
	loadRoutes(router)
	return router
}

// 新增加的路由文件需要在这里进行加载
func loadRoutes(router *gin.Engine) {
	router.Use(Middleware.Cors())
	// 注册请求API所需的路由
	api.LoadApi(router)
	admin.LoadAdmin(router)
	web.LoadWeb(router)

}
