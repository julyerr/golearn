package main

import (
	"fmt"
	"strconv"
)

func main() {
	//int 和 string 之间互相转换
	var i32v int32
	var i64v int64
	i32v = 123
	i64v = 123

	var iv int
	iv = int(i32v)
	s32 := strconv.Itoa(iv)
	iv = int(i64v)
	s64 := strconv.Itoa(iv)

	fmt.Println(strconv.ParseInt(s32,10,32))
	fmt.Println(strconv.ParseInt(s64,10,64))

	//float 和 string 之间互相转换
	var f32v float32
	var f64v float64
	f32v = 1.0
	f64v = 1.0
	sf32 := strconv.FormatFloat(float64(f32v),'f',-1,32)
	sf64 := strconv.FormatFloat(f64v,'f',-1,64)
	fmt.Println(sf32,sf64)

	v,_ := strconv.ParseFloat(sf32,32)
	f32v = float32(v)
	f64v,_ = strconv.ParseFloat(sf64,64)
	fmt.Println(f32v,f64v)
}
