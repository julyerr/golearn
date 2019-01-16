package main

import (
	"fmt"
	"time"
)

func worker(index int, exit chan bool) {
	<-exit
	fmt.Println("exit ", index)
}

func main() {
	//communication
	start := make(chan bool)
	for i := 0; i < 10; i++ {
		go worker(i, start)
	}
	close(start)
	time.Sleep(1 * time.Second)

	//trans data
	id:= newIdService()
	for i := 0; i < 10; i++ {
		fmt.Println(<- id)
	}

	//	select
}

func newIdService() chan string {
	id := make(chan string)
	go func() {
		var count int64
		for {
			id <- fmt.Sprintf("%x", count)
			count += 1
		}
	}()
	return id
}


