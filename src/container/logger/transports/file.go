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

	if utils.IsEnv(utils.Development, utils.Testing, utils.Staging) {
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

func (f File) Debug(args ...interface{}) {
	f.Context.Debugln(args...)
}

func (f File) Info(args ...interface{}) {
	f.Context.Infoln(args...)
}

func (f File) Warning(args ...interface{}) {
	f.Context.Warningln(args...)
}

func (f File) Error(args ...interface{}) {
	f.Context.Errorln(args...)
}
