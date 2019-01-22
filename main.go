package main

import (
	"errors"
	"fmt"
)

func main(){
	err := errors.New("error")
	fmt.Println(err)
}
