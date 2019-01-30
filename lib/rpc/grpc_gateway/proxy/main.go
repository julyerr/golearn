package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"golearn/lib/rpc/grpc_gateway/helloworld"
	"google.golang.org/grpc"
	"net/http"
)

var (
	flagEndpoint = flag.String("port",":9090","endpoint")
)

//curl -X POST localhost:8080/v1/example/echo -d '{"message":"ql"}'
func main(){
	flag.Parse()
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := helloWorld.RegisterHelloWorldHandlerFromEndpoint(ctx,mux,*flagEndpoint,opts)
	if err != nil{
		panic(err)
	}

	if err = http.ListenAndServe(":8080",mux);err != nil{
		panic(err)
	}
}
