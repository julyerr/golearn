package main

import "fmt"

type data struct{}

func (this *data) Error() string {
	return ""
}

func test() error {
	var p *data = nil
	//正确做法
	//if true{
	//	return nil
	//}
	return p
}

func main() {
	p := test()
	if p == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("Not nil")
	}
}
