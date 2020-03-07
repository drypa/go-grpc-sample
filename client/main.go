package main

import (
	"context"
	"github.com/drypa/go-grpc-sample/account"
	"google.golang.org/grpc"
	"log"
)

const address = "localhost:50000"

func main() {
	dial, err := grpc.Dial(address, grpc.WithInsecure())
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
