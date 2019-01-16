package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	type Test struct{
		Name string
		Passwd string
	}
	test := &Test{
		Name:"ql",
		Passwd:"ql",
	}
	o1,err := json.Marshal(test)
	if err != nil{
		panic(err)
	}
	o2,err := json.MarshalIndent(test,"","\t")
	if err != nil{
		panic(err)
	}
	fmt.Println(string(o1))
	fmt.Println(string(o2))
}