package main

import (
	"fmt"
	"log"
	"sync"
)

type syncMapTest struct {
	name string
	//nil map可以直接使用
	smap sync.Map
	ssmap map[string]sync.Map
	//但是origin map需要initial
	omap map[int32]string
}

func main() {
	smt := syncMapTest{}
	fmt.Println(smt.omap == nil)
	smt.smap.Store(1,1)
	tmpS := smt.ssmap["hello"]
	tmpS.Store(1,1)
	v, ok := tmpS.Load(1)
	if ok {
		log.Printf("load ok, v:%d", v.(int))
	}else {
		log.Printf("not found")
	}
}
