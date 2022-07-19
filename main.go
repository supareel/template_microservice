package main

import (
	"fmt"
	"gomicro/config"
	"gomicro/database"
	"gomicro/domain/taskmanager"
	"gomicro/pkg"
	"gomicro/proto/gen/proto"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func handleArgs(args []string, db database.Database) {
	if args[1] == "-migrate" {
		db.MigrateDB(40)
	}
}

func main() {
	godotenv.Load(".env.development")
	config.LoadEnvConfig()

	// connect to DB
	db, err := database.NewPostgresClient(config.EnvConfig.DB_USERNAME, config.EnvConfig.DB_PASS, config.EnvConfig.DB_SERVER, config.EnvConfig.DB_NAME, config.EnvConfig.DB_SSL, config.EnvConfig.DB_PORT, 30)
	if err != nil {
		pkg.FancyHandleError(err)
	}

	if len(os.Args) > 1 {
		handleArgs(os.Args, db)
		return
	}
	// repository
	repo := taskmanager.NewPostgresRepository(db)
	svc := taskmanager.NewTaskService(repo)
	taskRouter := taskmanager.NewTaskHandler(svc)
	// start server
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	proto.RegisterTaskServiceServer(grpcServer, taskRouter)

	listener, err := net.Listen("tcp", ":" + config.EnvConfig.APP_PORT)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Running gRPC Server on => 0.0.0.0:" + config.EnvConfig.APP_PORT)
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot Start Server : ", err)
	}
}
