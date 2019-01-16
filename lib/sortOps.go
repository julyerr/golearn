package main

import (
	"fmt"
	"sort"
)

type uL []*user1

type user1 struct{
	id int
	age int
	name string
}

func (u *user1) String() string{
	return fmt.Sprintf("id:%d,age:%d,name:%s",u.id,u.age,u.name)
}

func (ul uL) Len() int{
	return len(ul)
}

func (ul uL) Less(i,j int) bool{
	u1,u2 := ul[i],ul[j]
	if u1.id == u2.id{
		if u1.age ==u2.age{
			if u1.name ==  u2.name{
				return true
			}
			return u1.name < u2.name
		}
		return u1.age < u2.age
	}
	return u1.id < u2.id
}

func (ul uL) Swap(i,j int){
	u := &user1{}
	u = ul[i]
	ul[i] = ul[j]
	ul[j] = u
}


func main(){
	//内置的array和slice的排序方式
	sl := []int{
		89,23,2,0,1,190,-109,
	}
	fmt.Println(sl)
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] > sl[j]
	})
	fmt.Println(sl)


	//自定义排序，使用指针减少空间的复制代价
	var ul uL
	ul = []*user1{
		{id:1,age:1,name:"hello",},
		{id:2,age:1,name:"hello",},
		{id:1,age:2,name:"hello",},
		{id:1,age:1,name:"world",},
	}
	fmt.Println("before")
	for i:= range ul{
		fmt.Printf("%+v\n",ul[i])
	}

	fmt.Println("after")
	sort.Sort(ul)
	for i:= range ul{
		fmt.Printf("%+v\n",ul[i])
	}
}