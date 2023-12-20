package main

import (
	"golang-backend-microservice/config"
	"golang-backend-microservice/container/log"
	"golang-backend-microservice/dataservice"
	Nats "golang-backend-microservice/dataservice/nats"
	"golang-backend-microservice/model"
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

	// Mock Nats request
	params := model.MySqlReqArgs{
		Table: "book",
		Where: map[string]interface{}{"author": "JK Rowling"},
	}
	res, err := Nats.Request[model.MySqlReqArgs](nc, "database.sql.select", params)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Debug(string(res.Data))
	}

	runtime.Goexit()
}
