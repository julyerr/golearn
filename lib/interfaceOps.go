package main

import "fmt"

type Human interface{
	Say(string)
}

type userip struct{

}

func (u *userip) Say(name string){
	fmt.Println(name)
}

//接口对应的值发生了cp
func testip(u Human) Human{
	fmt.Printf("%x\n",&u)
	return u
}

func main(){
	fmt.Println()
	u := &userip{}
	fmt.Printf("%x\n",&u)
	uR := testip(u)
	fmt.Printf("%x\n",&uR)

	var human Human
	human = u
	fmt.Printf("%x\n",&human)
	//发生了值复制
	person := human.(*userip)
	fmt.Printf("%x\n",&person)

	BREAK:
	for i:=0;i<4;i++{
		switch human.(type) {
		case *userip:
			fmt.Println("userip type")
			break BREAK
		}
		fmt.Println("continue")
	}
}