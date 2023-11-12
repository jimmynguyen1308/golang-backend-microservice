package transport

import (
	"golang-backend-microservice/container/utils"
	"os"

	"github.com/rollbar/rollbar-go"
)

type Rollbar struct {
	Loglevel int32
	Alert    bool
}

const (
	ERROR   = 1
	WARNING = 2
	INFO    = 3
	DEBUG   = 4
)

func (r Rollbar) Default() Rollbar {
	token, exists := os.LookupEnv("ROLLBAR_ACCESS_TOKEN")
	if !exists || token == "" {
		r.Alert = false
	} else {
		r.Alert = true
		rollbar.SetToken(token)
	}

	rollbar.SetEnvironment(os.Getenv("ENVIRONMENT"))
	r.Loglevel = WARNING
	if utils.IsEnv(utils.Development, utils.Testing, utils.Staging) {
		r.Loglevel = DEBUG
	}

	return r
}

func (r Rollbar) Debug(args ...interface{}) {
	if r.Loglevel == DEBUG && r.Alert {
		rollbar.Debug(args...)
	}
}

func (r Rollbar) Info(args ...interface{}) {
	if r.Loglevel >= INFO && r.Alert {
		rollbar.Info(args...)
	}
}

func (r Rollbar) Warning(args ...interface{}) {
	if r.Loglevel >= WARNING && r.Alert {
		rollbar.Warning(args...)
	}
}

func (r Rollbar) Error(args ...interface{}) {
	if r.Loglevel >= ERROR && r.Alert {
		rollbar.Error(args...)
	}
}
