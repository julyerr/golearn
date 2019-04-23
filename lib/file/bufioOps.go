package main

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	file := "bufioOps.go"
	f,err := os.Open(file)
	if err != nil {
		logrus.Panicln(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		l,_,err := r.ReadLine()
		if err != nil {
			logrus.Printf("read failed:%s", err)
			return
		}
		logrus.Printf("read %s", l)
	}
}
