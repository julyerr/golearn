package time

import (
	"fmt"
	"time"
)

//https://blog.csdn.net/u011304970/article/details/72724357
func main(){
//	timer
	notify := make(chan bool,1)
	go func(time <-chan time.Time){
		<- time
		notify <- true
	}(time.After(1*time.Second))
	<- notify
	fmt.Println("main received notify")


	time.AfterFunc(1*time.Second,func (){
		fmt.Println("it's the time")
	})

	timer := time.NewTimer(1*time.Second)
	go func(){
		<-timer.C
		notify <- true
	}()
	<-notify
	fmt.Println("main received notify")


// 	循环定时器
	tick := time.NewTicker(1*time.Second)
	go func(){
		//tick stop之后，一直等待
		for t := range tick.C {
			fmt.Println(t)
		}
		fmt.Println("tick goroutine exist")
	}()
	time.Sleep(5*time.Second)
	tick.Stop()
	time.Sleep(2*time.Second)
}
