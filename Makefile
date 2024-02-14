mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(patsubst %/,%,$(dir $(mkfile_path)))

setup:
	./setup/setup
stoppg:
	docker stop postgres && docker rm postgres 

startpg:
	docker run --name postgres -p 4040:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root microservice_db

createtable:
	go run .\cmd\migrate\migrate.go

dropdb:
	docker exec -it postgres dropdb microservice_db

test:
	go test -v -cover ./internal/... ./database/... ./util/...

server:
	go run cmd/server/server.go

client:
	go run cmd/client/client.go

api-server:
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=$(current_dir)/server.gen.yaml $(current_dir)/task.yaml
api-client:
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=$(current_dir)/client.gen.yaml $(current_dir)/task.yaml
api-types:
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=$(current_dir)/types.gen.yaml $(current_dir)/task.yaml

.PHONY: setup stoppg startpg createdb dropdb test protogen server client entgen entinit