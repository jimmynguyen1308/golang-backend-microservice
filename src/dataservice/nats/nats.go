package nats

import (
	"fmt"
	"golang-backend-microservice/container/logger"
	"log"
	"net/http"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"github.com/rollbar/rollbar-go"
)

type ServiceConfig struct {
	ServiceName  string
	Version      string
	Description  string
	EndpointName string
}

func CreateNatsMicroservice() micro.Service {
	svcConfig := ServiceConfig{
		ServiceName:  "Database",
		Version:      "0.1.0-development",
		Description:  "Microservice for database requests and responses",
		EndpointName: "database",
	}

	logger.CreateTransports(logger.Console, logger.File, logger.Rollbar)
	rollbar.SetCodeVersion(svcConfig.Version)
	rollbar.SetCustom(map[string]interface{}{
		"ServiceName": svcConfig.ServiceName,
	})

	nc := openNatsConnection()
	svc := addNatsMicroService(svcConfig, nc)
	rollbar.SetCustom(map[string]interface{}{
		"ServiceName": svcConfig.ServiceName,
		"ServiceID":   svc.Info().ID,
	})
	return svc
}

func openNatsConnection() *nats.Conn {
	nc, err := nats.Connect(os.Getenv("NATS_HOST"),
		nats.UserInfo(os.Getenv("NATS_USER"), os.Getenv("NATS_PASS")))
	if err != nil {
		logger.Error("Error connecting to NATS: ", err)
		log.Fatalf("Error connecting to NATS: %s\n", err)
		return nil
	}
	return nc
}

func addNatsMicroService(config ServiceConfig, nc *nats.Conn) micro.Service {
	svc, err := micro.AddService(nc, micro.Config{
		Name:        config.ServiceName,
		Version:     config.Version,
		Description: config.Description,
		Endpoint: &micro.EndpointConfig{
			Subject: config.EndpointName,
			Handler: micro.HandlerFunc(func(req micro.Request) {
				req.Respond([]byte(fmt.Sprint(http.StatusOK)))
			}),
		},
		ErrorHandler: func(s micro.Service, n *micro.NATSError) {
			logger.Error(n.Error())
		},
	})
	if err != nil {
		logger.Error("Error adding service: ", err)
		log.Fatalf("Error adding service: %s\n", err)
		return nil
	}
	return svc
}
