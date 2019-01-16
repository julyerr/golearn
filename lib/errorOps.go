package main

import "fmt"

func main(){
	f1()
	fmt.Println("main")
}

func f1(){
	defer func(){
		if r := recover(); r != nil{
			fmt.Println("recover",r)
		}
	}()
	g1(0)
	fmt.Println("print in f1")
}

func g1(i int){
	if i > 3{
		fmt.Println("panic")
		panic(i)
	}
	defer fmt.Println("defer in g1",i)
	fmt.Println("print in g1",i)
	g1(i+1)
}
