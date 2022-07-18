package main

import (
	"context"
	"gomicro/internal/pb/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Cannot connect to server: ", err)
	}

	accountClient := proto.NewAccountServiceClient(conn)

	req := &proto.GetAllAccountsRequest{
		Val: 1,
	}
	res, err := accountClient.GetAllAccounts(context.Background(), req)
	if err != nil {
		log.Fatalln("Cannot connect to server: ", err)
	}

	log.Println(res)
}
