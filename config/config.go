package config

import (
	gin2 "github.com/gin-gonic/gin"
	"os"
	"strings"
	"time"
	_ "user_center/app"
	"user_center/pkg/env"
)

var (
	GinModel string
	Name     = os.Getenv("APP_NAME")
	URL      = os.Getenv("APP_URL")
	Env      = os.Getenv("APP_ENV")
	APPKey   = os.Getenv("APP_KEY")
	//Debug       = env.DefaultGetBool("DEBUG", false)
	Host        = os.Getenv("APP_HOST")
	CachePrefix = os.Getenv("CACHE_PREFIX")
	Database    = database{
		MySQL: map[string]MysqlConf{
			"default": {
				Host:        os.Getenv("DB_HOST"),
				Port:        os.Getenv("DB_PORT"),
				Username:    os.Getenv("DB_USERNAME"),
				Password:    os.Getenv("DB_PASSWORD"),
				Database:    os.Getenv("DB_DATABASE"),
				MaxLiftTime: time.Second * 60,
			},
		},
		Redis: map[string]RedisConf{
			"default": {
				Host:     env.DefaultGet("REDIS_HOST", "127.0.0.1").(string),
				Password: env.DefaultGet("REDIS_PASSWORD", "").(string),
				Port:     env.DefaultGetInt("REDIS_PORT", 6379),
				Database: env.DefaultGetInt("REDIS_DATABASE", 0),
			},
		},
	}
	Filesystems = filesystems{
		Default: "local",
		Cloud:   "",
		Disks: Disks{
			Local: Local{
				Driver: "local",
				Root:   "app/public",
			},
		},
	}
)

func init() {
	if !strings.EqualFold(Env, "local") &&
		!strings.EqualFold(Env, "production") &&
		!strings.EqualFold(Env, "testing") {
		panic("env APP_ENV must be: local, production, testing")
	}
	switch Env {
	case "testing":
		GinModel = gin2.TestMode
	case "local":
		GinModel = gin2.DebugMode
	case "production":
		GinModel = gin2.ReleaseMode
	}
}
