package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"path/filepath"
)

var (
	flagPath = flag.String("path", "", "path")
)

func main() {
	flag.Parse()

	fileInfo,err := os.Stat(*flagPath);
	if err != nil {
		log.Panicln("open %s failed, err:%s", *flagPath, err)
	}
	if !fileInfo.IsDir() {
		log.Panicf("path %s is not dir", *flagPath)
	}

	var count int
	err = filepath.Walk(*flagPath, func(path string, info os.FileInfo, err error) error {
		log.Printf("pathName :%s", path)
		count++
		if count == 2 {
			return errors.New("returnInfo")
		}
		return nil
	})
	if err != nil {
		if err.Error() == "returnInfo" {
			log.Println("return normal")
			return
		}
		log.Panicf("walk %s failed,err:%s", *flagPath, err)
	}
}
