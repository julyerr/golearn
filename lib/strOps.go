package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main(){
	str := "pre string"
	bufStr := bytes.NewBufferString(str)
	for i := 0; i < 3; i++  {
		bufStr.WriteString(randomString())
	}
	fmt.Println(bufStr.String())

	fmt.Println("replace str---------")
	str = "1,2"
	fmt.Println(strings.Replace(str,",","|",-1))
}

func randomString() string{
	return "randomStr"
}
