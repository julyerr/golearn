package sync

import (
	"fmt"
	"log"
	"sync"
)

type user2 struct{
	name string
	pw string
}

var (
	once = &sync.Once{}
)

func main(){
	for i:= 0 ;i < 2; i++ {
		once.Do(func() {
			log.Printf("'just do once")
		})
	}

	syncV := sync.Pool{
		New: func() interface{} {
			user := &user2{
				name:"hello",
				pw:"world",
			}
			return user
		},
	}

	u1 := syncV.Get().(*user2)
	fmt.Printf("%+v\n",u1)

	u2 := syncV.Get().(*user2)
	fmt.Printf("%+v\n",u2)

	syncV.Put(u1)
	u1 = syncV.Get().(*user2)
	fmt.Printf("%+v\n",u1)
}
