package main

import (
	"crypto/md5"
	"fmt"
)

func main(){
	md := md5.New()
	result := md.Sum([]byte("firstsecond"))
	fmt.Println(fmt.Sprintf("%x",result[:15]))
	md.Reset()
	md.Write([]byte("first"))
	md.Write([]byte("second"))
	fmt.Println(fmt.Sprintf("%x",result[:15]))
}
