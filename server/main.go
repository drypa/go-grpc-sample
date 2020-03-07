package main

import (
	"context"
	"github.com/drypa/go-grpc-sample/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	creds, err := credentials.NewServerTLSFromFile("../ssl/server.crt", "../ssl/server.key")
	if err != nil {
		log.Fatalf("failed to load TLS keys: %v", err)
	}
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	s := grpc.NewServer(opts...)
	server := server{}

	account.RegisterAccountServer(s, &server)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
