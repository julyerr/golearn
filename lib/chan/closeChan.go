package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
//https://www.jianshu.com/p/d24dfbb33781
func IsClosed(ch <-chan int) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}

func SafeSend(ch chan int, value int) (closed bool){
	defer func(){
		if recover() != nil{
			closed = true
		}
	}()
	ch <- value
	closed = false
	return
}

type MyChannel struct{
	ch chan int
	once sync.Once
}

func (my *MyChannel) NewChan() *MyChannel{
	return &MyChannel{ch:make(chan int,1)}
}

func (my *MyChannel) SafeClose(){
	my.once.Do(func(){
		close(my.ch)
	})
}

func main(){
	rand.Seed(time.Now().UnixNano())
	maxRandomNum := 100000
	workerNum := 100
	wg := sync.WaitGroup{}

	batchCh := make(chan int,100)

	//一个sender，多个receiver
	//go func(){
	//	for{
	//		if value:= rand.Intn(maxRandomNum); value == 0{
	//			close(batchCh)
	//		}else{
	//			batchCh <- value
	//		}
	//	}
	//}()
	//
	//for i := 0; i < workerNum; i++ {
	//	wg.Add(1)
	//	go func(){
	//		defer wg.Done()
	//		//关闭batchCh 之后for返回
	//		for value := range batchCh{
	//			fmt.Println(value)
	//		}
	//	}()
	//}

	//多个sender，一个receiver
	//stopCh := make(chan int)
	//for i := 0; i< workerNum;i++  {
	//	wg.Add(1)
	//	go func(){
	//		defer wg.Done()
	//
	//		for{
	//			value := rand.Intn(maxRandomNum)
	//			select {
	//				case <- stopCh:
	//					return
	//				case batchCh <- value:
	//			}
	//		}
	//	}()
	//}
	//
	//go func(){
	//	for value := range batchCh{
	//		if value ==  maxRandomNum - 1{
	//			close(stopCh)
	//		}
	//	}
	//}()

	//多个sender和receiver
	toStop := make(chan int)
	stopCh := make(chan int,1)

	//moderator
	go func(){
		<- toStop
		close(stopCh)
	}()

	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()

			for{
				// 是否退出
				select{
				case <- stopCh:
					return
				default:
				}

				value := rand.Intn(maxRandomNum)
				if value == 0{
					select{
					//塞值
					case toStop <- value:
					default:
					}
					return
				}

				select{
				case <- stopCh:
					return
				case batchCh <- value:
				}
			}
		}()
	}

	for i := 0; i < workerNum; i++ {
		wg.Add(1)

		go func(){
			defer wg.Done()

			select{
			case <- stopCh:
				return
			default:
			}

			select{
			case <- stopCh:
				return
			case  value := <- batchCh:
				if value == maxRandomNum - 1{
					select {
					case toStop <- value:
					default:
					}
					return
				}else{
					fmt.Println(value)
				}
			}
		}()
	}

	wg.Wait()
}