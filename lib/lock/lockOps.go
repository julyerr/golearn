package main

import (
	"fmt"
	"sync"
)

type testLock1 struct {
	name string
	sync.RWMutex
}

type testLock struct {
	sync.RWMutex
	name string
	test1 *testLock1
}

//锁之间只有RR之间兼容

func main(){
	test := &testLock{
		name: "test",
		test1: &testLock1{
			name: "test1",
		},
	}

	//goroutines

	//ch := make(chan int,1)
	//
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func(){
	//	defer wg.Done()
	//	time.Sleep(500*time.Millisecond)
	//	select{
	//	case <- time.After(1*time.Second):
	//		test.Lock()
	//		//test.name = "newtest1"
	//		test.test1.name = "newtest1"
	//		test.Unlock()
	//	case <-ch:
	//	}
	//	fmt.Println("go1 done")
	//}()
	//
	//test.RLock()
	//time.Sleep(2*time.Second)
	//test.RUnlock()
	//fmt.Println("unlock")
	//fmt.Println(test.test1.name)
	//wg.Wait()

	//self lock,可重入锁
	//test.RLock()
	//test.test1.Lock()
	//test.test1.Unlock()
	//test.RUnlock()

	test.Lock()
	test.test1.RLock()
	test.test1.RUnlock()
	test.Unlock()
	fmt.Println(test.name)
}
