stoppg:
	docker stop postgres && docker rm postgres 

startpg:
	docker run --name postgres -p 4040:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root microservice_db

dropdb:
	docker exec -it postgres dropdb microservice_db

test:
	go test -v -cover ./...

protogen:
	protoc --go_out=./proto/gen --go_opt=paths=source_relative --go-grpc_out=./proto/gen --go-grpc_opt=paths=source_relative proto/*.proto

server:
	go run cmd/server/main.go

client:
	go run cmd/client/main.go

entgen:
	go generate ./...

entinit:
	go run entgo.io/ent/cmd/ent init [ModelName]

.PHONY: stoppg startpg createdb dropdb test protogen server client entgen entinit