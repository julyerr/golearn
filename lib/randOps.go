package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(int64(time.Now().UnixNano()))
	fmt.Println(rand.Int31())
}
