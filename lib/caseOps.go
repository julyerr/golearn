package main

import (
	"fmt"
	"os"
)

func main(){
	var a int
	fmt.Fscanf(os.Stdin,"%d", &a)
	switch {
	case a == 1:
		//自动跳出
		fmt.Println("a")
	case a == 2:
		fmt.Println("b")
	case a >= 3 && a <= 4:
		fmt.Println("c")
	}

	BREAK:
	for i:=0;i<10;i++{
		if i == 5{
			break BREAK
		}
	}
}