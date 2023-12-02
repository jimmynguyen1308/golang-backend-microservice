package nats

import (
	"fmt"
	"golang-backend-microservice/container/log"
	"net/http"
	"time"

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

type Connection struct {
	User string
	Pass string
	Host string
	ServiceConfig
}

func (c Connection) Open() (*nats.Conn, micro.Service) {
	var nc *nats.Conn
	var svc micro.Service

	nc = c.openNatsConnection()
	if nc != nil {
		svc = c.addNatsService(nc)
		rollbar.SetCodeVersion(c.Version)
		rollbar.SetCustom(map[string]interface{}{
			"ServiceName": c.ServiceName,
			"ServiceID":   svc.Info().ID,
		})
		return nc, svc
	}

	return nc, svc
}

func (c Connection) openNatsConnection() *nats.Conn {
	nc, err := nats.Connect(
		c.Host, nats.UserInfo(c.User, c.Pass),
		nats.PingInterval(20*time.Second), nats.MaxPingsOutstanding(5),
	)
	if err != nil {
		log.Error(log.ErrNatsConnect, err)
		return nil
	}
	return nc
}

func (c Connection) addNatsService(nc *nats.Conn) micro.Service {
	svc, err := micro.AddService(nc, micro.Config{
		Name:        c.ServiceName,
		Version:     c.Version,
		Description: c.Description,
		Endpoint: &micro.EndpointConfig{
			Subject: c.EndpointName,
			Handler: micro.HandlerFunc(func(req micro.Request) {
				req.Respond([]byte(fmt.Sprint(http.StatusOK)))
			}),
		},
		ErrorHandler: func(s micro.Service, e *micro.NATSError) {
			log.Error(e.Error())
		},
	})
	if err != nil {
		log.Error(log.ErrNatsMicroAdd, err)
		return nil
	}
	return svc
}
