package main

import (
	"context"
	"golearn/lib/rpc/grpc/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":5001"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	return &helloworld.HelloResponse{Message: "from server:" + req.GetName()}, nil
}

func main() {
	ls, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(ls); err != nil {
		log.Fatalln(err)
	}
}
