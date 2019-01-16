package main

import (
	"flag"
	"os"
	"os/signal"
	"time"
)

var (
	flagConcurrency = flag.Int("con",10,"concurrency")
)

func main(){
	chans := make(chan int, *flagConcurrency)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		for i := 0; i < *flagConcurrency; i++ {
			chans <- 1
		}
		time.Sleep(1*time.Second)
		os.Exit(1)
	}()
}
