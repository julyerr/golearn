package main

import (
	"context"
	"fmt"
	"golearn/lib/rpc/grpc/helloworld"
	"google.golang.org/grpc"
	"time"
)

var (
	addr = "localhost:5001"
	name = "default"
)

func main(){
	conn,err := grpc.Dial(addr,grpc.WithInsecure())
	if err != nil{
		panic(err)
	}
	client := helloworld.NewGreeterClient(conn)
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()
	helloRequest := &helloworld.HelloRequest{Name:name}
	helloResp,err := client.SayHello(ctx,helloRequest)
	if err != nil{
		panic(err)
	}
	fmt.Println(helloResp.GetMessage())
}
