NAME=database
VERSION=0.1.0-development

all: help

.PHONY: help
help: Makefile
	@echo
	@echo " $(shell basename ${PWD}) 🎉"
	@echo
	@echo " Chose a command to run:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /' 
	@echo

## start: Run program in dev mode
.PHONY: start
start: test
	go run ./src/cmd/main.go

## test: Run unit tests
.PHONY: test
test:
	go clean -testcache
	go test -v ./src/container/utils/...
	go test -v ./src/dataservice/nats/...
	go test -v ./src/dataservice/mysql/...

## build: Generate Docker image
.PHONY: build
build:
	go build -o ./build/${NAME} ./src/cmd
	docker build -t ${NAME}:${VERSION} ./
