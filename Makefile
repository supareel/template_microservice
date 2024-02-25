mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(patsubst %/,%,$(dir $(mkfile_path)))

setup:
	./setup/setup

openapi-gen:
	swag init -g ../../cmd/server/server.go -d ./src/internal/taskmanager

start:
	docker compose up

stop:
	docker compose down

.PHONY: setup start stop