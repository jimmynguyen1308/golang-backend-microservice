package main

import (
	"golang-backend-microservice/config"
	"golang-backend-microservice/container/logger"
	"golang-backend-microservice/dataservice/mysql"
	"golang-backend-microservice/dataservice/nats"
	"log"

	"runtime"

	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go/micro"
)

func main() {
	config.Init()

	svc := nats.CreateNatsMicroservice()
	mysql := mysql.OpenMySqlConnection()
	if svc != nil && mysql != nil {
		if err := mockSql(svc, mysql); err != nil {
			logger.Error(err)
			log.Fatal(err)
		}
	}

	runtime.Goexit()
}

func mockSql(svc micro.Service, mysql *sqlx.DB) error {
	return nil
}
