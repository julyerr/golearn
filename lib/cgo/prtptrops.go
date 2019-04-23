package main

/*
#include <stdlib.h>

int* x;

void test(void** in){
	int i = 0 ;
	x = malloc(10*sizeof(int));
	for (i = 0;i < 10 ;i++) {
		x[i] = i;
	}
	in = &x;
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main(){
	var data unsafe.Pointer
	C.test(&data)
	fmt.Printf("%+v", data)

	data1 := unsafe.Pointer(C.malloc(8))
	C.test(&data1)
	fmt.Printf("%+v", data1)
}