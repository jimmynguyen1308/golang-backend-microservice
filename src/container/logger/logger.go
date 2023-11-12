package logger

import transport "golang-backend-microservice/container/logger/transports"

const (
	Console = "console"
	File    = "file"
	Rollbar = "rollbar"
)

type Transports interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
}

var destinations []Transports

func CreateTransports(transports ...string) {
	for _, t := range transports {
		switch t {
		case Console:
			destinations = append(destinations, transport.Console{}.Default())
		case File:
			destinations = append(destinations, transport.File{}.Default())
		case Rollbar:
			destinations = append(destinations, transport.Rollbar{}.Default())
		}
	}
}

func Debug(args ...interface{}) {
	for _, t := range destinations {
		t.Debug(args...)
	}
}

func Info(args ...interface{}) {
	for _, t := range destinations {
		t.Info(args...)
	}
}

func Warning(args ...interface{}) {
	for _, t := range destinations {
		t.Warning(args...)
	}
}

func Error(args ...interface{}) {
	for _, t := range destinations {
		t.Error(args...)
	}
}
