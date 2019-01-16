package main

import (
	"bytes"
	"fmt"
)

func main(){
	strs := "hello world"
	rBuf := bytes.NewBufferString(strs)

	sl := make([]byte,8)
	rBuf.Read(sl)
	fmt.Printf("%s\n",sl)
	rBuf.Read(sl)
	fmt.Printf("%s\n",sl)

	wBuf := bytes.NewBufferString("")
	wBuf.WriteString("hello world\n")
	fmt.Printf(wBuf.String())
	wBuf.Reset()
	fmt.Println(wBuf.String())
}
