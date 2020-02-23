package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "server/account"
)

type server struct {
	pb.UnimplementedAccountServer
}

const address = ":50000"

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAccountServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
