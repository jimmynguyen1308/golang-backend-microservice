package transport

import (
	"golang-backend-microservice/container/utils"
	"os"

	log "github.com/sirupsen/logrus"
)

type File struct {
	Context *log.Entry
}

func (f File) Default() File {
	loglevel := log.ErrorLevel

	if utils.IsEnv(utils.ENV_DEVELOPMENT, utils.ENV_TESTING, utils.ENV_STAGING) {
		loglevel = log.DebugLevel
	}

	file, err := os.OpenFile("out.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error("unable to open log file")
		return f
	}
	f.Context = log.NewEntry(&log.Logger{
		Out:       file,
		Formatter: &log.JSONFormatter{},
		Level:     loglevel,
		Hooks:     make(log.LevelHooks),
	})

	return f
}

func (f File) Debug(format string, args ...interface{}) {
	f.Context.Debugf(format, args...)
}

func (f File) Info(format string, args ...interface{}) {
	f.Context.Infof(format, args...)
}

func (f File) Warn(format string, args ...interface{}) {
	f.Context.Warningf(format, args...)
}

func (f File) Error(format string, args ...interface{}) {
	f.Context.Errorf(format, args...)
}
