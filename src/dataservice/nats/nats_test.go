package nats_test

import (
	"encoding/json"
	"fmt"
	"golang-backend-microservice/container/log"
	"golang-backend-microservice/dataservice/nats"
	"golang-backend-microservice/model"
	"testing"

	"github.com/nats-io/nats-server/v2/server"
	mockserver "github.com/nats-io/nats-server/v2/test"
	"github.com/nats-io/nats.go/micro"
)

var (
	NATS_USER string = "local"
	NATS_PASS string = "password"
	TEST_PORT int    = 8369

	structRes1 = nats.StatusResponse{
		Status: 200,
	}
	expectedRes1 string = "{\"status\":200}"

	structRes2 = nats.StatusResponse{
		Status: 400,
		Error:  "Invalid request",
	}
	expectedRes2 string = "{\"status\":400,\"error\":\"Invalid request\"}"

	structRes3 = nats.DataResponse[model.Book]{
		Status: 200,
		Data:   []model.Book{},
	}
	expectedRes3 string = "{\"status\":200,\"data\":[]}"
)

func mockNatsServer(port int) *server.Server {
	options := mockserver.DefaultTestOptions
	options.Port = port
	options.Username = NATS_USER
	options.Password = NATS_PASS
	return mockserver.RunServer(&options)
}

func mockResponse() micro.HandlerFunc {
	return func(req micro.Request) {
		var d model.MySqlReqArgs
		json.Unmarshal(req.Data(), &d)
		switch d.Table {
		case "1":
			structRes1.Respond(req)
			return
		case "2":
			structRes2.Respond(req)
			return
		case "3":
			structRes3.Respond(req)
			return
		default:
			return
		}
	}
}

func TestNatsConnection(t *testing.T) {
	server := mockNatsServer(TEST_PORT)
	defer server.Shutdown()

	// Test open NATS connection
	nc, svc := nats.Connection{
		User: NATS_USER,
		Pass: NATS_PASS,
		Host: fmt.Sprintf("nats://127.0.0.1:%d", TEST_PORT),
		ServiceConfig: nats.ServiceConfig{
			ServiceName:  "Database",
			Version:      "1.0.0",
			Description:  "Lorem ipsum",
			EndpointName: "database",
		},
	}.Open()
	defer nc.Close()
	if nc == nil || svc == nil {
		t.Errorf("Error: Nats connection unsuccessful")
	}

	// Test NATS connection request & respond functions
	mock := svc.AddGroup("mock")
	{
		if err := mock.AddEndpoint("respond", mockResponse()); err != nil {
			log.Error(log.ErrNatsMicroAdd, err.Error())
		}
	}
	config := model.MySqlReqArgs{
		Table: "1",
		Where: map[string]interface{}{},
	}
	res1, err := nats.Request[model.MySqlReqArgs](nc, "mock.respond", config)
	if err != nil {
		t.Errorf("Error: Nats request failed - %v", err)
	} else if string(res1.Data) != expectedRes1 {
		t.Errorf("Expected %s, got %s: Incorrect response", expectedRes1, string(res1.Data))
	}
	config.Table = "2"
	res2, err := nats.Request[model.MySqlReqArgs](nc, "mock.respond", config)
	if err != nil {
		t.Errorf("Error: Nats request failed - %v", err)
	} else if string(res2.Data) != expectedRes2 {
		t.Errorf("Expected %s, got %s: Incorrect response", expectedRes2, string(res2.Data))
	}
	config.Table = "3"
	res3, err := nats.Request[model.MySqlReqArgs](nc, "mock.respond", config)
	if err != nil {
		t.Errorf("Error: Nats request failed - %v", err)
	} else if string(res3.Data) != expectedRes3 {
		t.Errorf("Expected %s, got %s: Incorrect response", expectedRes3, string(res3.Data))
	}
}

func BenchmarkOpenNatsConnection(b *testing.B) {
	server := mockNatsServer(TEST_PORT)
	defer server.Shutdown()

	for n := 0; n < b.N; n++ {
		nc, svc := nats.Connection{
			User: NATS_USER,
			Pass: NATS_PASS,
			Host: fmt.Sprintf("nats://127.0.0.1:%d", TEST_PORT),
			ServiceConfig: nats.ServiceConfig{
				ServiceName:  "Database",
				Version:      "1.0.0",
				Description:  "Lorem ipsum",
				EndpointName: "database",
			},
		}.Open()
		defer func() {
			nc.Close()
			svc.Stop()
		}()
	}
}
