package logger

import (
	"io"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/lierbai/nspider/core/common/config"
	"github.com/lierbai/nspider/core/common/formatter"
	"github.com/lierbai/nspider/core/common/util"
	log "github.com/sirupsen/logrus"
)

func init() {
	conf := config.ConfigI
	if util.IsDirExists(conf.LogDir) {
		if err := os.MkdirAll(conf.LogDir, 0755); err != nil {
			panic(err)
		}
	}
	logfile := path.Join(conf.LogDir, conf.Name)
	fsWriter, err := rotatelogs.New(
		logfile+conf.LogSuffix,
		rotatelogs.WithMaxAge(conf.LogMaxTime*time.Hour),
		rotatelogs.WithRotationTime(conf.LogMaxTime*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(fsWriter, os.Stdout)
	log.SetReportCaller(true)
	log.SetFormatter(formatter.NewLogFormatter())
	log.SetOutput(multiWriter)
	log.SetLevel(log.InfoLevel)
}
