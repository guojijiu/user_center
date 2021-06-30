package glog

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	conf "user_center/config"
	"user_center/pkg/glog/hook"
)

// 默认日志格式
var defaultFormat = LocalFormatter{&log.JSONFormatter{
	PrettyPrint:       true,
	DisableHTMLEscape: true,
}}
var Channels = map[string]*log.Logger{}
var openFiles []*os.File

type LocalFormatter struct {
	log.Formatter
}

// 会一次性载入所有通道
func Init() {
	loadChannels()
}

func (u LocalFormatter) Format(e *log.Entry) ([]byte, error) {
	e.Time = e.Time.Local()
	return u.Formatter.Format(e)
}

func Default() *log.Logger {
	return Channel("default")
}

func Channel(name string) *log.Logger {
	if l, ok := Channels[name]; ok {
		return l
	}
	if c, ok := conf.Logging.Channels[name]; ok {
		Channels[name] = configLog(c)
		return Channels[name]
	}
	return Channels["default"]
}

// 加载所有通道
func loadChannels() {
	// default
	Channels["default"] = configDefaultLog()
	// channels
	for name, logConf := range conf.Logging.Channels {
		Channels[name] = configLog(logConf)
	}
}

func configDefaultLog() *log.Logger {
	if logC, ok := conf.Logging.Channels[conf.Logging.Default]; ok {
		l := log.StandardLogger()
		config(l, logC)
		return l
	}
	return nil
}

func configLog(logConf conf.Log) *log.Logger {
	l := log.New()
	config(l, logConf)
	return l
}

func config(l *log.Logger, c conf.Log) {
	var err error
	var format log.Formatter
	if c.LogFormatter != nil {
		format = c.LogFormatter
	} else {
		format = defaultFormat
	}
	l.SetLevel(c.Level)
	l.SetReportCaller(true)
	// add hooks
	for _, h := range c.Hooks {
		l.AddHook(h)
	}
	if c.Driver == conf.Daily {
		// 日驱动
		l.AddHook(hook.NewLfsHook(c.Path, 7, format))
	} else {
		l.SetFormatter(format)
		err = os.MkdirAll(filepath.Dir(c.Path), os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("Create dir %s failed: [%+v]", filepath.Dir(c.Path), err))
		}
		f, err := os.OpenFile(c.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Sprintf("Create file %s failed: [%+v]", c.Path, err))
		}
		openFiles = append(openFiles, f)
		l.SetOutput(f)
	}
}

func Close() {
	for _, f := range openFiles {
		_ = f.Close()
	}
}

func SetLogFiels(f *map[string]interface{}) log.Fields {
	fields := log.Fields{}
	if f != nil {
		for k, v := range *f {
			fields[k] = v
		}
	}
	return fields
}
