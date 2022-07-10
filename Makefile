set_env:
	@export GOPATH=$HOME
	@export GO111MODULE=on
	@go mod tidy

run:
	@go run cmd/main.go

build:
	@go build cmd/main.go