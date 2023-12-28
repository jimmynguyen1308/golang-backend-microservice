package main

import (
	"golang-backend-microservice/config"
	"golang-backend-microservice/container/log"
	"golang-backend-microservice/dataservice"
	Nats "golang-backend-microservice/dataservice/nats"
	"golang-backend-microservice/model"
	"os"
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

	r := dataservice.SetupRoutes(nc)
	r.Run(":" + os.Getenv("SERVER_PORT"))

	// Mock Nats request
	params := model.MySqlReqArgs{
		Table: "book",
		Where: map[string]interface{}{"author": "JK Rowling"},
	}

	// TODO:
	// - [x] Add Gin routes to /container
	// - [ ] Bind each route to each NATS request accordingly
	// - [x] Add function setupRoutes() to main.go

	res, err := Nats.Request[model.MySqlReqArgs](nc, "database.sql.select", params)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Debug(string(res.Data))
	}

	runtime.Goexit()
}
