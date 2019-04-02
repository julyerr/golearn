package main

import "log"

func main() {
	preT,preH := 10,10
	tmpT,tmpH := 11,11
	tmp :=float64(tmpH-preH)/float64(tmpT-preT)
	log.Println(tmp)
	log.Printf("hint rate %f", tmp)
}
