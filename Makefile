.PHONY: ;
.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

BINARY_NAME='user-service'

build: build-go-server

build-go-server:
	rm -f $(BINARY_NAME)
	go build -o $(BINARY_NAME) cmd/question_keeper_service/question_service.go


run-app:
	go run cmd/question_keeper_service/question_service.go start

start: run-app



