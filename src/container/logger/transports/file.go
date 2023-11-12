package transport

import log "github.com/sirupsen/logrus"

type File struct {
	Context *log.Entry
}

func (f File) Default() File {
	return f
}

func (f File) Debug(args ...interface{}) {

}

func (f File) Info(args ...interface{}) {

}

func (f File) Warning(args ...interface{}) {

}

func (f File) Error(args ...interface{}) {

}
