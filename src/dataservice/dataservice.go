package dataservice

import (
	"golang-backend-microservice/config"
	"golang-backend-microservice/container/log"
	"golang-backend-microservice/container/utils"
	MongoDb "golang-backend-microservice/dataservice/mongodb"
	MySql "golang-backend-microservice/dataservice/mysql"
	Nats "golang-backend-microservice/dataservice/nats"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
)

func OpenNatsConnection() *nats.Conn {
	version := "0.1.0"
	if !utils.IsEnv(utils.ENV_PRODUCTION) {
		version += "-" + os.Getenv("ENVIRONMENT")
	}

	nc, svc := Nats.Connection{
		User: os.Getenv("NATS_USER"),
		Pass: os.Getenv("NATS_PASS"),
		Host: os.Getenv("NATS_HOST"),
		ServiceConfig: Nats.ServiceConfig{
			ServiceName:  "Database",
			Version:      version,
			Description:  "Microservice for database requests and responses",
			EndpointName: "database",
		},
	}.Open()
	if nc == nil || svc == nil {
		return nil
	}

	var sql *sqlx.DB
	for {
		sql = MySql.Connection{
			User: os.Getenv("MYSQL_USER"),
			Pass: os.Getenv("MYSQL_PASS"),
			Host: os.Getenv("MYSQL_HOST"),
		}.Open()
		if sql != nil {
			break
		}
		log.Debug(log.DebugConnectionRetry, config.RETRY_TIMER)
		time.Sleep(config.RETRY_TIMER * time.Second)
	}

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
		mongo := database.AddGroup("mongo")
		if err := mongo.AddEndpoint("find", MongoDb.FindRecord()); err != nil {
			log.Error(log.ErrNatsMicroAdd, err.Error())
			return nc
		}
		if err := mongo.AddEndpoint("insert", MongoDb.InsertRecord()); err != nil {
			log.Error(log.ErrNatsMicroAdd, err.Error())
			return nc
		}
		if err := mongo.AddEndpoint("update", MongoDb.UpdateRecord()); err != nil {
			log.Error(log.ErrNatsMicroAdd, err.Error())
			return nc
		}
		if err := mongo.AddEndpoint("delete", MongoDb.DeleteRecord()); err != nil {
			log.Error(log.ErrNatsMicroAdd, err.Error())
			return nc
		}
	}

	return nc
}
