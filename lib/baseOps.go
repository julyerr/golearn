package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
)

func main(){
	data,err := ioutil.ReadFile("mock_iot.bin")
	if err != nil {
		log.Panicln(err)
	}
	s := base64.StdEncoding.EncodeToString(data)
	log.Printf(s)
	//feature :=  "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"
	//data,err := base64.StdEncoding.DecodeString(feature)
	//if err != nil {
	//	log.Panicln(err)
	//}
	//log.Printf("size feature:%d", len(data))
	//data,err := ioutil.ReadFile("../tmp/image/test.jpg")
	//if err != nil{
	//	panic(err)
	//}
	//str := base64.StdEncoding.EncodeToString(data)
	//type imageWrapper struct{
	//	Content string
	//}
	//iw := &imageWrapper{
	//	Content: str,
	//}
	//data1,err := json.Marshal(iw)
	//if err != nil{
	//	panic(err)
	//}
	//
	//var iw2 imageWrapper
	//err = json.Unmarshal(data1,&iw2)
	//if err != nil{
	//	panic(err)
	//}
	//
	//ioutil.WriteFile("../tmp/image/test_new1.jpg",[]byte(iw2.Content),0644)
	//data,err = base64.StdEncoding.DecodeString(iw2.Content)
	//if err != nil{
	//	panic(err)
	//}
	//ioutil.WriteFile("../tmp/image/test_new2.jpg", data,0644)

}
