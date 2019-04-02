package main

import (
	"flag"
	"fmt"
)

func init() {
	fmt.Println("init func")
	fmt.Println(*flagTest)
}

var (
	flagTest = flag.String("test", "test", "test")
)

func main() {
	flag.Parse()
}
