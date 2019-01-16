package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os/exec"
	"strconv"
	"time"
)

var (
	flagPid = flag.Int("pid", 0, "pid")
	flagCputimes = flag.Int("cpu",3,"cpu times")
	maxCpu float64
)

const (
	CPU_DELTA=0.2
	CPU_COUNT = 1
)

func main() {

	flag.Parse()
	if *flagPid == 0 {
		fmt.Println("--pid")
		return
	}
	fmt.Println(captureCpuInfo(*flagPid))
}

func staticsCpuInfo() (result string,err error){
	var sum float64
	var max float64
	for i:=0; i<*flagCputimes; i++ {
		s,e := captureCpuInfo(*flagPid)
		if e != nil{
			return "",e
		}
		max = math.Max(max,s)
		if max < maxCpu{
			i--
		}
		time.Sleep(500*time.Millisecond)
	}
	maxCpu = max
	fmt.Println("max",max)
	for i:=0; i<*flagCputimes; i++ {
		s,e := captureCpuInfo(*flagPid)
		if e != nil{
			return "",e
		}
		//找到比较准的
		if math.Abs(max-s) < CPU_DELTA*max{
			sum += s
		}else{
			i--
		}
		time.Sleep(500*time.Millisecond)
	}
	return strconv.FormatFloat(sum/float64(*flagCputimes*CPU_COUNT),'f',-1,64),nil
}

func captureCpuInfo(pid int) (float64,error){
	spid := strconv.Itoa(pid)
	//top -b -n 1 -p $pid  2>&1 | awk -v pid=$pid '{if ($1 == pid)print $9}'
	cmd1 := exec.Command("top", "-b", "-n", "1", "-p", spid)
	//cmd2 := exec.Command("awk","-v","pid="+spid,"{if ($1 == "+spid+")print $9}")
	cmd2 := exec.Command("grep",spid)
	cmd2.Stdin, _ = cmd1.StdoutPipe()
	cmd3 := exec.Command("awk","{print $9}")
	cmd3.Stdin, _ = cmd2.StdoutPipe()
	stdout, _ := cmd3.StdoutPipe()
	//cmd2.Stdout = os.Stdout
	//cmd2.Stderr = os.Stderr

	err := cmd3.Start()
	if err != nil {
		fmt.Println("cmd3 start", err.Error())
		return 0,err
	}
	err = cmd2.Start()
	if err != nil {
		fmt.Println("cmd2 start", err.Error())
		return 0,err
	}
	err = cmd1.Run()
	if err != nil {
		fmt.Println("cmd1 run", err.Error())
		return 0,err
	}
	//err = cmd2.Wait()
	//if err != nil{
	//	fmt.Println("cmd2 wait",err.Error())
	//	return
	//}
	fmt.Println("before read")
	outBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("reader err", err.Error())
		return 0, err
	}
	result := string(outBytes[:len(outBytes)-1])
	//strs := strings.Split(line," ")
	//if len(strs) < 17 {
	//	return 0,errors.New("top result failed")
	//}
	//有时候读出返回0，尝试
	r,err := strconv.ParseFloat(result,64)
	if err != nil {
		if  result == "" {
			time.Sleep(1*time.Second)
			return captureCpuInfo(pid)
		}
		return 0,err
	}
	return r,nil
}