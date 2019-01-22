package main

import "fmt"

type Human interface{
	Say(string)
}

type userip struct{
	ip string
}

func (u *userip) Say(name string){
	fmt.Println(name)
}

func main(){
	fmt.Println()
	u := &userip{
		ip: "name",
	}
	fmt.Printf("%x\n",&u.ip)

	//接口只是引用，实际地址值没有发生变化
	var human1 Human
	human1 = u
	fmt.Printf("%x\n",&human1.(*userip).ip)

	BREAK:
	for i:=0;i<4;i++ {
		switch human1.(type) {
		case *userip:
			fmt.Println("userip type")
			break BREAK
		}
		fmt.Println("continue")
	}
}