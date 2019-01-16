package main

import "fmt"

type notifier interface {
	notify1()
	notify2()
}

type user struct {
	name  string
	email string
}

func (u *user) notify1() {
}

func (u *user) notify2(){
}

func (u *user) test() {
	fmt.Println("test")
}

func (u user) test1() {
	fmt.Println("test1")
}

func main() {
	u := &user{
		name:  "zhangsan",
		email: "zhangsan",
	}

	uSer := user{name: "zhangsan", email: "test"}
	sendNotification(u)
	// user没有实现接口
	sendNotification(uSer)
}

func sendNotification(n notifier) {
	n.notify1()
}
