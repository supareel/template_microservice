package main

import (
	"fmt"
	"gomicro/config"
	"gomicro/database"
	"gomicro/domain/taskmanager"
	"gomicro/pkg"
	proto "gomicro/pkg/proto"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func main() {
	godotenv.Load(".env.development")
	config.LoadEnvConfig()

	// connect to DB
	db, err := database.NewPostgresClient(config.EnvConfig.DB_USERNAME, config.EnvConfig.DB_PASS, config.EnvConfig.DB_SERVER, config.EnvConfig.DB_NAME, config.EnvConfig.DB_SSL, config.EnvConfig.DB_PORT, 30)
	if err != nil {
		pkg.FancyHandleError(err)
	}
	defer db.CloseClient()
	// repository
	repo := taskmanager.NewPostgresRepository(db)
	svc := taskmanager.NewTaskService(repo)
	taskRouter := taskmanager.NewTaskHandler(svc)
	// start server
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	proto.RegisterTaskServiceServer(grpcServer, taskRouter)
	endpoint := fmt.Sprintf(":%d",config.EnvConfig.APP_PORT)
	listener, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatal(err)
	}

	// run grpc
	fmt.Printf("Running gRPC Server on => %s\n", endpoint)
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot Start Server : ", err)
	}
}