package main

import (
	"golang-backend-microservice/config"
	"golang-backend-microservice/container/log"
	"golang-backend-microservice/dataservice"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	config.Init()

	var nc *nats.Conn
	for {
		nc = dataservice.OpenNatsConnection()
		if nc != nil {
			break
		}
		log.Debug(log.DebugConnectionRetry, config.RETRY_TIMER)
		time.Sleep(config.RETRY_TIMER * time.Second)
	}
	log.Info(log.InfoConnectionCreated)
	runtime.Goexit()
}
