package transport

import (
	"golang-backend-microservice/container/utils"
	"os"

	log "github.com/sirupsen/logrus"
)

type Console struct {
	Context *log.Entry
}

func (c Console) Default() Console {
	formatter := log.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	loglevel := log.ErrorLevel

	if utils.IsEnv(utils.Development, utils.Testing, utils.Staging) {
		formatter = log.TextFormatter{
			DisableColors: false,
			FullTimestamp: false,
		}
		loglevel = log.DebugLevel
	}

	c.Context = log.NewEntry(&log.Logger{
		Out:       os.Stderr,
		Formatter: &formatter,
		Level:     loglevel,
		Hooks:     make(log.LevelHooks),
	})

	return c
}

func (c Console) Debug(args ...interface{}) {
	c.Context.Debugln(args...)
}

func (c Console) Info(args ...interface{}) {
	c.Context.Infoln(args...)
}

func (c Console) Warning(args ...interface{}) {
	c.Context.Warningln(args...)
}

func (c Console) Error(args ...interface{}) {
	c.Context.Errorln(args...)
}
