package hook

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
	"time"
)

func BasenameNotSuffix(filename string) string {
	if len(filename) == 0 {
		return filename
	}
	suffix := filepath.Ext(filename)
	return strings.TrimSuffix(filename, suffix)
}

// 拆分日志文件
func NewLfsHook(logName string, maxRemainCnt uint, defaultLogFormat log.Formatter) log.Hook {
	writer, err := rotatelogs.New(
		BasenameNotSuffix(logName)+"_%Y%m%d%H"+filepath.Ext(logName),
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(logName),
		// WithRotationTime设置日志分割的时间，这里设置为24小时分割一次
		rotatelogs.WithRotationTime(time.Hour*24),
		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(maxRemainCnt)),
	)

	if err != nil {
		log.Errorf("config local file system for logger error: %v", err)
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, defaultLogFormat)

	return lfsHook
}
