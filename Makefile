postgres:
	sudo docker stop postgres && docker rm postgres && docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	sudo docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	sudo docker exec -it postgres dropdb simple_bank

test:
	go test -v -cover ./...

protogen:
	rm -rf ./proto/*.pb.go && protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative proto/*.proto

server:
	go run cmd/server/main.go

client:
	go run cmd/client/main.go

.PHONY: postgres createdb dropdb test protogen server client