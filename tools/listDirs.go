package tools

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	listDirs(dir, 0)
}

func listDirs(dir string, level int) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}
	for _, fileInfo := range fileInfos {
		fmt.Print("|")
		for i := 0; i < level; i++ {
			fmt.Print(" ")
		}
		fmt.Print("-")
		if fileInfo.IsDir() {
			fmt.Println(fileInfo.Name() + string(os.PathSeparator))
			listDirs(dir+string(os.PathSeparator)+fileInfo.Name(), level+1)
		} else {
			fmt.Println(fileInfo.Name())
		}
	}
}
