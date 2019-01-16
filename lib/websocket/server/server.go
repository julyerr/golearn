package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)


var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	flagPort = flag.String("port", "9003", "listen port")
)

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	_,message,err := conn.ReadMessage()
	if err != nil{
		return
	}
	fmt.Println("server receive",string(message))
	if err != nil{
		return
	}
	for{
		err = conn.WriteMessage(websocket.TextMessage,[]byte("hello world"))
		if err != nil{
			//fmt.Println("server send failed",err)
			return
		}
	}
}


func main() {
	flag.Parse()

	http.HandleFunc("/test", handler)
	panic(http.ListenAndServe("0.0.0.0:"+*flagPort, nil))
}

