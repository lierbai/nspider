package logger

import (
	"io"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/lierbai/nspider/core/common/config"
	"github.com/lierbai/nspider/core/common/util"
	log "github.com/sirupsen/logrus"
)

func init() {
	conf := config.ConfigI
	if util.IsDirExists(conf.LogDir) != true {
		if err := os.MkdirAll(conf.LogDir, 0755); err != nil {
			panic(err)
		}
	}
	logfile := path.Join(conf.LogDir, conf.Name+"_%Y-%m-%d."+conf.LogSuffix)
	fsWriter, err := rotatelogs.New(
		logfile,
		rotatelogs.WithMaxAge(conf.LogMaxTime*time.Hour),
		rotatelogs.WithRotationTime(conf.LogRotationTime*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(fsWriter, os.Stdout)
	log.SetReportCaller(true)
	log.SetFormatter(NewLogFormatter())
	log.SetOutput(multiWriter)
	log.SetLevel(log.DebugLevel)
	log.Debug("======== Spider Logging init ========")
}
