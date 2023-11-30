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

	if utils.IsEnv(utils.ENV_DEVELOPMENT, utils.ENV_TESTING, utils.ENV_STAGING) {
		formatter = log.TextFormatter{
			DisableColors:   false,
			FullTimestamp:   true,
			TimestampFormat: "15:04:05.000000",
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

func (c Console) Debug(format string, args ...interface{}) {
	c.Context.Debugf(format, args...)
}

func (c Console) Info(format string, args ...interface{}) {
	c.Context.Infof(format, args...)
}

func (c Console) Warn(format string, args ...interface{}) {
	c.Context.Warningf(format, args...)
}

func (c Console) Error(format string, args ...interface{}) {
	c.Context.Errorf(format, args...)
}
