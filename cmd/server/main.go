package main

import (
	"fmt"
	"gomicro/database"
	"gomicro/pb/proto"
	"gomicro/service"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main(){
  godotenv.Load(".env.development")
	database.ConnectToDB()

  accountsServer := service.NewAccountsService()
  grpcServer := grpc.NewServer()

  proto.RegisterAccountServiceServer(grpcServer, accountsServer)
  fmt.Println("Running gRPC Server on => 0.0.0.0:8080")

  listener, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatal(err)
  }
  err = grpcServer.Serve(listener)
  if err != nil {
    log.Fatal("Cannot Start Server : ", err)
  }
}