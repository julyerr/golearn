package os

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main(){
	defer func() {
		fmt.Println("test defer exec after exit")
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("receive interrupt")
		os.Exit(1)
	}()
	time.Sleep(10*time.Second)
}
