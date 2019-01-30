package main

import (
	"context"
	"flag"
	"golearn/lib/rpc/grpc_gateway/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var (
	flagPort = flag.String("port",":9090","grpc server port")
)

type server struct{}

func (s *server) SayHello(ctx context.Context,req *helloWorld.StringMessage) (*helloWorld.StringMessage, error){
	return &helloWorld.StringMessage{
		Message: "Hello "+req.GetMessage(),
	},nil
}

func main(){
	flag.Parse()
	ls,err := net.Listen("tcp",*flagPort)
	if err != nil{
		panic(err)
	}
	s := grpc.NewServer()
	helloWorld.RegisterHelloWorldServer(s,&server{})
	reflection.Register(s)
	if err := s.Serve(ls);err != nil{
		panic(err)
	}
}
