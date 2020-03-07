package main

import (
	"context"
	"github.com/drypa/go-grpc-sample/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const address = "localhost:50000"

func main() {
	creds, err := credentials.NewClientTLSFromFile("../ssl/server.crt", "")
	if err != nil {
		log.Fatalf("failed to load TLS cert: %v", err)
	}

	dial, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Printf(err.Error())
		return
	}
	client := account.NewAccountClient(dial)
	ctx := context.Background()
	request := account.RegisterRequest{
		TelegramId: 0,
	}
	response, err := client.Register(ctx, &request)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	log.Printf(response.Id)

}
