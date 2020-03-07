package main

import (
	"context"
	"github.com/drypa/go-grpc-sample/account"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	account.UnimplementedAccountServer
}

func (s *server) Register(context.Context, *account.RegisterRequest) (*account.RegisterReply, error) {
	return &account.RegisterReply{
		TelegramId: 0,
		Id:         "123",
	}, nil
}

const address = ":50000"

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := server{}
	account.RegisterAccountServer(s, &server)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
