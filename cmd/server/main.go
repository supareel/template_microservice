package main

import (
	"fmt"
	"gomicro/controller"
	"gomicro/database"
	"gomicro/pb/proto"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load(".env.development")
	database.ConnectToDB()

	accountsServer := controller.NewAccountsService()
	grpcServer := grpc.NewServer()

	proto.RegisterAccountServiceServer(grpcServer, accountsServer)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Running gRPC Server on => 0.0.0.0:8080")
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot Start Server : ", err)
	}
}
