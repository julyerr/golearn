package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"math/rand"
	"time"
)

func main(){
	//rateLimit :=  rate.NewLimiter(2,4)
	//ctx,cancel := context.WithCancel(context.Background())
	//fmt.Println("in main",time.Now())
	//for i:=0;i<50;i++{
	//	go func(i int){
	//		rateLimit.Wait(ctx)
	//		fmt.Println(i,time.Now())
	//	}(i)
	//}
	//time.Sleep(5*time.Second)
	//cancel()
	//fmt.Println("in main",time.Now())

	rateLimit :=  rate.NewLimiter(2,4)
	fmt.Println("in main",time.Now())
	for i:=0;i<50;i++{
		go func(i int){
			time.Sleep(time.Duration(rand.Intn(10000))*time.Millisecond)
			if rateLimit.Allow(){
				fmt.Println(i,time.Now())
			}
		}(i)
	}
	time.Sleep(10*time.Second)
	fmt.Println("in main",time.Now())
}
