package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	"strconv"
)

func main() {
	var dialer *websocket.Dialer
	u := url.URL{Scheme: "ws", Host: "localhost:" + strconv.Itoa(9003), Path: "/test"}
	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.WriteMessage(websocket.TextMessage,[]byte("connect to server"))
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return
		}
		fmt.Println("client received",string(data))
	}
}
