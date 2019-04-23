package main

import (
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	defer func() {
		wg.Wait()
	}()

	wg.Add(1)
	go func() {
		time.Sleep(10*time.Second)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		defer func() {
			logrus.Printf("panic")
			wg.Done()

		}()
		time.Sleep(1*time.Second)
		panic(1)
	}()
}
