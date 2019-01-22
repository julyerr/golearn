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
	////communication
	//start := make(chan bool)
	//for i := 0; i < 10; i++ {
	//	go worker(i, start)
	//}
	//close(start)
	//time.Sleep(1 * time.Second)
	//
	////trans data
	//id:= newIdService()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<- id)
	//}
	//
	////	select

	//pubsub
	//chans := make(chan int,2)
	//mk := make(map[string]chan int)
	//mk["id1"]=chans
	//mk["id2"]=chans
	//tmps := make(chan int,2)
	//var i int
	//go func(){
	//	for{
	//		i++
	//		time.Sleep(time.Second)
	//		select{
	//		case tmps <- i:
	//		default:
	//		}
	//	}
	//}()
	//for{
	//	tmp := <-tmps
	//	for id,ch := range mk{
	//		select{
	//		case ch <- tmp:
	//			fmt.Println(id,tmp)
	//			<-ch
	//		default:
	//		}
	//	}
	//}

	//difference
	content := make(chan []byte,1)
	//content := make(chan []byte)
	go func(){
		var i int
		for{
			select{
			case content <- []byte("hello"):
			default:
			}
			time.Sleep(1*time.Second)
			i++
			if i ==  3{
				close(content)
				break
			}
		}
	}()
	BREAK:
	for {
		select{
		case val := <- content:
			if len(val) == 0{
				fmt.Println("exit")
				break BREAK
			}
			fmt.Println(string(val))
		default:
		}
	}
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


