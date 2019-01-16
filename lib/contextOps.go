package main

import (
	"context"
	"fmt"
	"time"
)

//值传递
func func1(ctx context.Context) {
	ctx = context.WithValue(ctx, "k1", "v1")
	func2(ctx)
}

func func2(ctx context.Context) {
	fmt.Println(ctx.Value("k1"))
}

func main() {
	ctx := context.Background()
	func1(ctx)

	//	超时等控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	func3(ctx)
}

func func3(ctx context.Context) {
	resp := make(chan bool, 1)
	go func() {
		time.Sleep(3 * time.Second)
		resp <- true
	}()
	select {
	case <-resp:
		fmt.Println("data received")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
	return
}
