package main

import (
	"context"
	"fmt"
	"gomicro/config"
	"gomicro/database"
	"gomicro/internal/taskmanager"
	"gomicro/pkg"
	proto "gomicro/pkg/proto"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	grpcEndpoint := fmt.Sprintf(":%d", config.EnvConfig.GRPC_PORT)

	
	// run grpc
	go func() {
		proto.RegisterTaskServiceServer(grpcServer, taskRouter)
		listener, err := net.Listen("tcp", grpcEndpoint)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Running gRPC Server on => %s\n", grpcEndpoint)
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatal("Cannot Start GRPC Server : ", err)
		}
	}()
	
	// run http
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	conn, err := grpc.DialContext(
		context.Background(),
		grpcEndpoint,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	proto.RegisterTaskServiceHandler(ctx, mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	httpEndpoint := fmt.Sprintf(":%d", config.EnvConfig.HTTP_PORT)
	log.Printf("Serving gRPC-Gateway on http://0.0.0.0%s\n", httpEndpoint)
	if err = http.ListenAndServe(httpEndpoint, mux); err != nil {
		log.Fatal("Cannot Start HTTP Proxy Server : ", err)
	}

}
