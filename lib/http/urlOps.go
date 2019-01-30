package main

import (
	"fmt"
	"net/url"
)

func main(){
	link := "http://www.baidu.com/page?name=ql&pw=ql&city=中国#frame1"
	encodedUrl := url.QueryEscape(link)
	fmt.Println("encoded",encodedUrl)
	decodedUrl,err := url.QueryUnescape(encodedUrl)
	if err != nil{
		panic("decoded failed"+err.Error())
	}
	fmt.Println("decoded",decodedUrl)

	url1,err := url.Parse(link)
	if err != nil{
		panic(err)
	}
	fmt.Printf("host:%s path:%s rawpath:%s rawquery:%s\n",url1.Host,url1.Path,url1.RawPath,
		url1.RawQuery)
	values,err := url.ParseQuery(link)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%+v\n",values)
}
