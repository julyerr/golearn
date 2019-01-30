package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
)

func main(){

	data,err := ioutil.ReadFile("../tmp/image/test.jpg")
	if err != nil{
		panic(err)
	}
	str := base64.StdEncoding.EncodeToString(data)
	type imageWrapper struct{
		Content string
	}
	iw := &imageWrapper{
		Content: str,
	}
	data1,err := json.Marshal(iw)
	if err != nil{
		panic(err)
	}

	var iw2 imageWrapper
	err = json.Unmarshal(data1,&iw2)
	if err != nil{
		panic(err)
	}

	ioutil.WriteFile("../tmp/image/test_new1.jpg",[]byte(iw2.Content),0644)
	data,err = base64.StdEncoding.DecodeString(iw2.Content)
	if err != nil{
		panic(err)
	}
	ioutil.WriteFile("../tmp/image/test_new2.jpg", data,0644)

}
