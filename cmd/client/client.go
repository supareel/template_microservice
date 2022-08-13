package main

import (
	"context"
	"fmt"
	"log"

	proto "gomicro/pkg/proto"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	"google.golang.org/grpc"
)

func main() {
	connectTo := "127.0.0.1:7005"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		  fmt.Errorf("failed to connect to PetStoreService on %s: %w", connectTo, err)
		return
	}
	log.Println("Connected to", connectTo)

	taskService := proto.NewTaskServiceClient(conn)
	resp, err := taskService.GetAllTasks(context.Background(), &proto.GetAllTasksRequest{})
	if err != nil {
		fmt.Errorf("failed to PutPet: %w", err)
		return
	}

	fmt.Printf("DATA ::\n%+v\nSTATUS ::\n%+v", resp.Data, resp.Status)
}
