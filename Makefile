#!make
include .env
# export $(shell 's/=.*//' .env)
export $(shell sed '/^\#/d; s/=.*//' .env) 
export APP=mcabank-auth

compose:
	docker-compose -f ./docker/docker-compose.yaml up -d

run:
	go run cmd/main.go
	
test:
	go test ./...

deploy:
	go build -o $(APP) cmd/main.go
	./$(APP) > $(APP).log 2>&1 &

logs:
	clear && tail -n 60 -f $(APP).log

