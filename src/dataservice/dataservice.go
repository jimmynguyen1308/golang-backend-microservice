package dataservice

import (
	"golang-backend-microservice/container/log"
	MySql "golang-backend-microservice/dataservice/mysql"
	Nats "golang-backend-microservice/dataservice/nats"
	"os"

	"github.com/nats-io/nats.go"
)

func OpenNatsConnection() *nats.Conn {
	nc, svc := Nats.Connection{
		User: os.Getenv("NATS_USER"),
		Pass: os.Getenv("NATS_PASS"),
		Host: os.Getenv("NATS_HOST"),
		ServiceConfig: Nats.ServiceConfig{
			ServiceName:  "Database",
			Version:      "0.1.0-development",
			Description:  "Microservice for database requests and responses",
			EndpointName: "database",
		},
	}.Open()
	if nc == nil || svc == nil {
		return nil
	}

	sql := MySql.Connection{
		User: os.Getenv("MYSQL_USER"),
		Pass: os.Getenv("MYSQL_PASS"),
		Host: os.Getenv("MYSQL_HOST"),
	}.Open()

	database := svc.AddGroup("database")
	{
		// MySQL endpoints
		mysql := database.AddGroup("sql")
		if err := mysql.AddEndpoint("select", MySql.SelectRecord(sql)); err != nil {
			log.Error(log.ErrNatsMicroAdd, err.Error())
			return nc
		}
		if err := mysql.AddEndpoint("insert", MySql.InsertRecord(sql)); err != nil {
			log.Error(log.ErrNatsMicroAdd, err.Error())
			return nc
		}
		if err := mysql.AddEndpoint("update", MySql.UpdateRecord(sql)); err != nil {
			log.Error(log.ErrNatsMicroAdd, err.Error())
			return nc
		}
		if err := mysql.AddEndpoint("delete", MySql.DeleteRecord(sql)); err != nil {
			log.Error(log.ErrNatsMicroAdd, err.Error())
			return nc
		}

		// MongoDB endpoints
		// mongo := database.AddGroup("mongo")
	}

	return nc
}
