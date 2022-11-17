package boot

import (
	"log"
	"user_center/app"
	"user_center/app/Http/Middleware"
	validator2 "user_center/app/validator"
	"user_center/pkg/db"
	"user_center/pkg/glog"
	migrate2 "user_center/pkg/migrate"
	"user_center/pkg/redis"
	"user_center/pkg/storage"
	"user_center/routes"
)

func SetInTest() {
	app.InTest = true
}

func SetInConsole() {
	app.InConsole = true
}

// 应用启动入口
func Boot() {
	var err error
	glog.Init()
	storage.Init(app.StoragePath)

	if _, err = redis.InitDef(); err != nil {
		log.Panicf("Init Default Redis connection filed: %+v", err)
	}

	if _, err = db.InitDef(); err != nil {
		log.Panicf("Init Default MySQL connection filed: %+v", err)
	}

	// 命令行模式下不加载路由
	if !app.InConsole {
		// 注册中间件
		Middleware.Init()
		// 注册路由
		router := routes.InitRouter()
		app.SetEngineRouter(router)
	}
	// 初始化自定义验证类
	validator2.Init()

	app.Booted = true

	migrate()
}

func Destroy() {
	glog.Close()
	redis.Close()
}

func migrate() {
	_ = db.Def().AutoMigrate(&migrate2.Migration{})
}
