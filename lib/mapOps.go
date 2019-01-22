package main

import "fmt"

type user3 struct{
	name string
}

func main(){
	u1 := &user3{name:"ql"}
	u2 := &user3{name:"ql"}
	//u3 := &user3{name:"ql1"}
	map1 := make(map[user3]string)
	map1[*u1]="test"
	if val,ok := map1[*u2];ok{
		fmt.Println("exists",val)
	}else{
		fmt.Println("no exists")
	}

}
