NAME=database
VERSION=0.1.0-development

dev:
	go run ./src/cmd/main.go

test:
	go clean -testcache
	go test -v ./src/container/utils/...
	go test -v ./src/dataservice/nats/...
	go test -v ./src/dataservice/mysql/...

build:
	GOOS=linux GOARCH=arm64 go build -o ./build/${NAME} ./src/cmd
	docker build -t ${NAME}:${VERSION} ./

clean:
	docker rmi -f $$(docker images -f "dangling=true" -q)