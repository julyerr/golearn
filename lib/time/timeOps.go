package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	t,err := time.Parse("2006-01-02 15:04:05","2006-01-02 15:04:05")
	if err != nil{
		panic(err)
	}
	fmt.Println(t)
}
