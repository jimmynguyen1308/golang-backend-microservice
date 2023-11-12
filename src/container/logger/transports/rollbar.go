package transport

import log "github.com/sirupsen/logrus"

type Rollbar struct {
	Context *log.Entry
}

func (r Rollbar) Default() Rollbar {
	return r
}

func (r Rollbar) Debug(args ...interface{}) {

}

func (r Rollbar) Info(args ...interface{}) {

}

func (r Rollbar) Warning(args ...interface{}) {

}

func (r Rollbar) Error(args ...interface{}) {

}
