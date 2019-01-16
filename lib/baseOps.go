package main

import (
	"encoding/base64"
	"io/ioutil"
)

func main(){
	//str := "HAAAAAAAAAAAABIAGAAUAAAADAAAAAgAAAAEABIAAAAUAAAAIAAAAKA3XBJoAQAALAAAAAoAAABwaG90b19wYXRoAAAHAAAAY2FwdHVyZQAYAAAAAAAAAAQABAAEAAAACAAAAAAAAAA="
	//src,_ := base64.StdEncoding.DecodeString(str)
	//fmt.Printf("%x\n",src)

	data,err := ioutil.ReadFile("test.jpg")
	if err != nil{
		panic(err)
	}
	str := base64.StdEncoding.EncodeToString(data)
	ioutil.WriteFile("test_new.jpg",[]byte(str),0644)
}
