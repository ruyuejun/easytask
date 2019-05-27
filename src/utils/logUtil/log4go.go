package logUtil

import (
	"github.com/sirupsen/logrus"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"time"
	"path"
	"github.com/pkg/errors"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.SetLevel(logrus.InfoLevel)
	ConfigLocalFilesystemLogger("./logs", "log", time.Hour*24, time.Second*20)
}

func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPaht := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d%H%M%S",
		//rotatelogs.WithLinkName(baseLogPaht),      	// 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             		// 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), 		// 日志切割时间间隔
	)

	if err != nil {
		Log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	wm := lfshook.WriterMap {
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}

	tf := logrus.TextFormatter {
		DisableColors: true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	}

	lfHook := lfshook.NewHook( wm, &tf)

	Log.AddHook(lfHook)
}
