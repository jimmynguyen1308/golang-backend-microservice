package nats_test

import (
	"fmt"
	"golang-backend-microservice/dataservice/nats"
	"testing"

	"github.com/nats-io/nats-server/v2/server"
	mockserver "github.com/nats-io/nats-server/v2/test"
)

const (
	NATS_USER string = "local"
	NATS_PASS string = "password"
	TEST_PORT int    = 8369
)

func MockNatsServer(port int) *server.Server {
	options := mockserver.DefaultTestOptions
	options.Port = port
	options.Username = NATS_USER
	options.Password = NATS_PASS
	return mockserver.RunServer(&options)
}

func TestOpenNatsConnection(t *testing.T) {
	server := MockNatsServer(TEST_PORT)
	defer server.Shutdown()

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
}

func BenchmarkOpenNatsConnection(b *testing.B) {
	server := MockNatsServer(TEST_PORT)
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
