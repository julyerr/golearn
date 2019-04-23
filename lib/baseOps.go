package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const str = "hello"

func main(){
	//bstr := base64.StdEncoding.EncodeToString([]byte(str))
	//maps := make(map[string]interface{})
	//maps["hello"] = bstr
	//data,err := json.Marshal(maps)
	//if err != nil {
	//	panic(err)
	//}
	//ioutil.WriteFile("base.json", data, 0644)

	data,err := ioutil.ReadFile("base.json")
	if err != nil {
		panic(err)
	}
	maps := make(map[string][]byte)
	err = json.Unmarshal(data, &maps)
	if err != nil {
		panic(err)
	}
	strb := maps["hello"]
	fmt.Println(string(strb))
}
