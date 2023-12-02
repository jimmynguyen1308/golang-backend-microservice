package main

import (
	"golang-backend-microservice/config"
	"golang-backend-microservice/container/log"
	"golang-backend-microservice/dataservice"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
)

const RETRY_TIMER = 30

func main() {
	config.Init()

	var nc *nats.Conn
	for {
		nc = dataservice.OpenNatsConnection()
		if nc == nil {
			log.Debug("Retry in %d seconds...", RETRY_TIMER)
			time.Sleep(RETRY_TIMER * time.Second)
		} else {
			break
		}
	}
	log.Info(log.InfoNatsMicroCreated)
	runtime.Goexit()
}
