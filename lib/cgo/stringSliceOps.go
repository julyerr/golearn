package main

//https://studygolang.com/articles/16812
//https://tonybai.com/2012/09/26/interoperability-between-go-and-c/

/*
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
char* cstr = "hello,world";
int cArray[] = {1,2,3,4};
char buf[256];

void print(char* str){
	printf("%s\n",str);
}
char* cat(char* str1,char* str2){
	strcpy(buf,str1);
	strcat(buf,str2);
	return buf;
}

void fill_255(char *buf,int32_t len){
	int32_t i;
	for (i = 0;i<len;i++){
		buf[i] = (char)255;
	}
}
 */
import "C"

import (
	"fmt"
	"unsafe"
)

func main(){
	//string
	str1,str2:= C.CString("hello,"),C.CString("world")
	defer C.free(unsafe.Pointer(str1))
	defer C.free(unsafe.Pointer(str2))
	str3 := C.GoString(C.cat(str1,str2))
	fmt.Println(str3)

	b := make([]byte,5)
	fmt.Println(b)
	//slice,array
	C.fill_255((*C.char)(unsafe.Pointer(&b[0])),C.int32_t(len(b)))
	fmt.Println(b)
	fmt.Println(cArrayToGoArray(unsafe.Pointer(&C.cArray[0]),4))
	fmt.Printf("%d\n",&C.cArray[0])
	fmt.Println(C.cArray[0])
//	针对c频繁访问数组空间的情况，使用C.malloc直接在c中分配不需要跨越访问


}

func cArrayToGoArray(cArray unsafe.Pointer,size int) (goarray []int){
	//p := uintptr(cArray)
	//for i:=0;i<size;i++{
	//	j := *(*int)(unsafe.Pointer(p))
	//	goarray = append(goarray,j)
	//	p += unsafe.Sizeof(j)
	//}
	//return

	return (*[1<<30]int)(cArray)[0:size]
}