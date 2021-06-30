package config

import (
	log "github.com/sirupsen/logrus"
	"user_center/app"
	"user_center/pkg/env"
	"user_center/pkg/glog/hook"
)

const (
	Daily = "daily" // 日驱动
	//Single = "single" // 单文件驱动
)

type logs map[string]Log

type Log struct {
	Driver       string
	Path         string
	Level        log.Level
	Days         int
	LogFormatter log.Formatter
	Hooks        []log.Hook
}

var(
	Logging = struct {
		Channels logs
		Default  string
	}{
		Channels: logs{
			"default": Log{
				Driver: env.DefGetStr("LOG_DEFAULT_DRIVER", Daily),
				Path:   app.GetStoragePath("logs/def/main.log"),
				Level:  log.DebugLevel,
				Days:   7,
				Hooks: []log.Hook{
					&hook.DefaultFieldHook{
						AppName: Name,
						AppUrl:  URL,
						AppEnv:  Env,
					},
				},
			},
			"gin": Log{
				Driver:       env.DefGetStr("LOG_DEFAULT_DRIVER", Daily),
				Path:         app.GetStoragePath("logs/route/route.log"),
				Level:        log.DebugLevel,
				Days:         7,
				LogFormatter: &log.TextFormatter{},
			},
			"db": Log{
				Driver: env.DefGetStr("LOG_DEFAULT_DRIVER", Daily),
				Path:   app.GetStoragePath("logs/db/db.log"),
				Level:  log.DebugLevel,
				Days:   7,
			},
			"request": Log{
				Driver: env.DefGetStr("LOG_DEFAULT_DRIVER", Daily),
				Path:   app.GetStoragePath("logs/request/request.log"),
				Level:  log.DebugLevel,
				Days:   30,
			},

		},
		Default: "default",
	}
)