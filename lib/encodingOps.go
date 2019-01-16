package main

import (
	"encoding/base64"
	"fmt"
)

func main(){
	toSend := []byte("toSend")
	s := base64.StdEncoding.EncodeToString(toSend)
	fmt.Println(s)
	toReceive,_ := base64.StdEncoding.DecodeString(s)
	fmt.Println(string(toReceive))
}
