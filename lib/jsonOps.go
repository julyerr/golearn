package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	type Test struct{
		Name string
		Passwd []byte
	}
	//map -> struct
	maps := make(map[string]interface{})
	maps["name"] = "name"
	maps["passwd"] = []byte("pw")
	data,err := json.Marshal(maps)
	if err != nil{
		panic(err)
	}
	var t Test
	json.Unmarshal(data,&t)
	fmt.Printf("%+v\n",t)

	//indent show
	test := &Test{
		Name:"ql",
		Passwd: []byte("ql"),
	}
	//o1,err := json.Marshal(&test)
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

	//nil show
	test1 := &Test{
		Name:"ql",
	}
	o3,err := json.Marshal(&test1)
	if err != nil{
		panic(err)
	}
	var test2 Test
	err = json.Unmarshal([]byte(o3),&test2)
	if err != nil{
		panic(err)
	}
	fmt.Println(test2.Passwd == nil)
}