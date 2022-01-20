postgres:
	sudo docker stop postgres && docker rm postgres && docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	sudo docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	sudo docker exec -it postgres dropdb simple_bank

test:
	go test -v -cover ./...

protogen:
	rm -rf ./proto/*.pb.go && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto

.PHONY: postgres createdb dropdb test protogen