package main

import (
	"bytes"
	"fmt"
)

func main(){
	str := "pre string"
	bufStr := bytes.NewBufferString(str)
	for i := 0; i < 3; i++  {
		bufStr.WriteString(randomString())
	}
	fmt.Println(bufStr.String())
}

func randomString() string{
	return "randomStr"
}
