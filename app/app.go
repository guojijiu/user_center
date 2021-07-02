package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"runtime"
)

// 存放公用变量, 用以解决循环依赖问题

var (
	EngineRouter *gin.Engine
	InTest       = false
	InConsole    = false
	AppPath      = ""
	TestPath     = ""
	StoragePath  = ""
	DatabasePath = ""
	Booted       = false
)

func init() {
	appPath := getBasePath()
	AppPath = appPath
	TestPath = fmt.Sprintf("%s/tests", appPath)
	StoragePath = fmt.Sprintf("%s/storage", appPath)
	DatabasePath = fmt.Sprintf("%s/database", appPath)
}

func SetEngineRouter(engine *gin.Engine) {
	EngineRouter = engine
}

func GetEngineRouter() *gin.Engine {
	return EngineRouter
}

func GetStoragePath(path string) string {
	return filepath.Join(StoragePath, path)
}

func GetDatabasePath(path string) string {
	return filepath.Join(DatabasePath, path)
}

// 获取项目基础路径的绝对值
func getBasePath() string {
	_, b, _, _ := runtime.Caller(1)
	return filepath.Join(filepath.Dir(b), "../")
}
