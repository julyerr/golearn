package main

import (
	"fmt"
	"regexp"
)

func main(){
	isValidBucketName := regexp.MustCompile(`^[A-Za-z0-9-]+$`).MatchString
	if isValidBucketName("photoalbumql"){
		fmt.Println("valid")
	}else{
		fmt.Println("invalid")
	}
}
