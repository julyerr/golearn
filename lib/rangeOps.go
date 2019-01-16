package main

import (
	"fmt"
	"time"
)

func main() {
	strs := []string{"hello", "world"}
	//v地址不变，输出之前值变成world
	for _, v := range strs {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(time.Second)

	array := []int{1,2,3}
	maps := make(map[int]*int)
	for k,v := range array{
		maps[k] = &v
	}
	for _,v := range maps{
		fmt.Println(*v)
	}

	for k,v := range array{
		value := v
		maps[k] = &value
	}
	for _,v := range maps{
		fmt.Println(*v)
	}
}
