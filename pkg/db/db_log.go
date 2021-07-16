package db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"time"
	"user_center/pkg/glog"
)

var newLogger logger.Interface

type Writer struct {
}

func (w Writer) Printf(_ string, args ...interface{}) {
	glog.Channel("db").WithFields(log.Fields{
		"sql_path": args[0],
		"use_time":      fmt.Sprintf("%vms", args[1]),
		"rows":          fmt.Sprintf("%v", args[2]),
		"sql":           args[3],
	}).Info("sql请求")
}

func initSqlLog() {
	newLogger = logger.New(
		Writer{},
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // 慢sql
			LogLevel:                  logger.Info,            // sql登记
			IgnoreRecordNotFoundError: true,                   // 过滤查找不到的错误
			Colorful:                  true,                   // 是否使用颜色
		},
	)
}
