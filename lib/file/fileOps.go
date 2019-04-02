package file

import (
	"fmt"
	"io/ioutil"
)

type user struct{
	name string
	pw string
}

func main(){
	//file,err := os.OpenFile("test.txt",os.O_CREATE|os.O_RDWR|os.O_APPEND,0644)
	//defer file.Close()
	//if err != nil{
	//	panic(err)
	//}
	//输入输出中写入文件
	//input := os.Stdin
	//defer input.Close()
	//var name string
	//var pw string
	//fmt.Fscanf(input,"%s %s",&name,&pw)
	//user := &user{
	//	name:name,
	//	pw:pw,
	//}
	//fmt.Printf("%+v\n",user)
	//fmt.Fprintf(file,"%s\t%s\n",user.name,user.pw)

	//读取文件
	//buf := make([]byte,1024)
	//for{
	//	n,err := file.Read(buf)
	//	if err != nil{
	//		break
	//	}
	//	fmt.Printf("%s",buf[:n])
	//}

	//bufio 拷贝文件
	//wFile,err := os.OpenFile("test.txt.cp",os.O_RDWR|os.O_CREATE|os.O_TRUNC,0644)
	//defer wFile.Close()
	//if err != nil{
	//	panic(err)
	//}
	//oi := bufio.NewReader(file)
	//oo := bufio.NewWriter(wFile)
	//for{
	//	string,err := oi.ReadString('\n')
	//	if err != nil{
	//		break
	//	}
	//	fmt.Println("read line",string)
	//	oo.WriteString(string)
	//}
	////文件内容需要flush
	//oo.Flush()

	dir := "/Users/qiulai/work/go/src/github.com/julyerr/golearn/lib"
	fs,err  := ioutil.ReadDir(dir)
	if err != nil{
		panic(err)
	}
	for i:=0 ;i< len(fs);i++{
		fmt.Println(fs[i].Name() )
	}
}
