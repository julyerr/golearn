package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

//http://www.01happy.com/golang-http-client-get-and-post/
func main(){
	//Get
	resp,err := http.Get("http://www.baidu.com")
	if err != nil{
		panic(err)
	}
	data,err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil{
		panic(err)
	}
	fmt.Println(len(data))

	//Post
	//maps := make(map[string]interface{})
	//maps["name"]="name"
	//data,err = json.Marshal(maps)
	//if err != nil{
	//	panic(err)
	//}
	//_,err = http.Post("http://www.baidu.com/page?name=ql","application/json",bytes.NewReader(data))
	//if err != nil{
	//	panic(err)
	//}

	req,err := http.NewRequest("GET","http://www.baidu.com",nil)
	req.Header.Set("k1","v1")
	req = req.WithContext(context.Background())
	client := &http.Client{}
	_,err = client.Do(req)
	if err != nil{
		panic(err)
	}
}
