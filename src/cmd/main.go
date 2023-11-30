package main

import (
	"golang-backend-microservice/config"
	"golang-backend-microservice/container/log"
	"golang-backend-microservice/dataservice"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {
	config.Init()

	nc := dataservice.OpenNatsConnection()
	if nc != nil {
		if err := mockSql(nc); err != nil {
			log.Error("Error: %s", err)
			return
		}
	}
	log.Info(log.InfoNatsMicroCreated)

	runtime.Goexit()
}

func mockSql(nc *nats.Conn) error {
	return nil
}
