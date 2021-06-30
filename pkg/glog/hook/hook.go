package hook

import (
	log "github.com/sirupsen/logrus"
)

// 为日志字段中添加一些系统描述
type DefaultFieldHook struct {
	AppName string
	AppUrl  string
	AppEnv  string
}

func (hook *DefaultFieldHook) Fire(entry *log.Entry) error {
	entry.Data["app_name"] = hook.AppName
	entry.Data["app_url"] = hook.AppUrl
	entry.Data["app_env"] = hook.AppEnv
	return nil
}

func (hook *DefaultFieldHook) Levels() []log.Level {
	return log.AllLevels
}
