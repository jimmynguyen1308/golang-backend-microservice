package transport

import (
	"fmt"
	"golang-backend-microservice/container/utils"
	"os"

	"github.com/rollbar/rollbar-go"
)

type Rollbar struct {
	Loglevel int32
	Alert    bool
}

const (
	errors  = 1
	warning = 2
	info    = 3
	debug   = 4
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
	r.Loglevel = warning
	if utils.IsEnv(utils.ENV_DEVELOPMENT, utils.ENV_TESTING, utils.ENV_STAGING) {
		r.Loglevel = debug
	}

	return r
}

func (r Rollbar) Debug(format string, args ...interface{}) {
	if r.Loglevel == debug && r.Alert {
		rollbar.Debug(fmt.Sprintf(format, args...))
	}
}

func (r Rollbar) Info(format string, args ...interface{}) {
	if r.Loglevel >= info && r.Alert {
		rollbar.Info(fmt.Sprintf(format, args...))
	}
}

func (r Rollbar) Warn(format string, args ...interface{}) {
	if r.Loglevel >= warning && r.Alert {
		rollbar.Warning(fmt.Sprintf(format, args...))
	}
}

func (r Rollbar) Error(format string, args ...interface{}) {
	if r.Loglevel >= errors && r.Alert {
		rollbar.Error(fmt.Sprintf(format, args...))
	}
}
