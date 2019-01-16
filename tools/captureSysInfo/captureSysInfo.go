package captureSysInfo

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	CPU_COUNT = 56
	STATIC_TIMES = 3
)

const (
	TYPE_CPU = iota
	TYPE_MEM
	TYPE_C_M
)

func process(exit chan int, staticType int, pid int) ([]string, error) {
	//do something with goroutine
	time.Sleep(5 * time.Second)
	s, e := staticsInfo(staticType, pid)
	if e != nil {
		return nil, e
	}
	time.Sleep(300 * time.Millisecond)
	return s, nil
}

//取中位数，然后重复三次求平均值
func staticsInfo(staticType int, pid int) (result []string, err error) {
	var sum1, sum2 float64
	for i := 0; i < 3; i++ {
		sl := make([][]float64, 2)
		for i := 0; i < STATIC_TIMES; i++ {
			s, e := captureInfo(pid, staticType)
			if e != nil {
				return nil, e
			}
			if staticType != TYPE_C_M {
				sl[0] = append(sl[0], s[0])
			} else {
				sl[0] = append(sl[0], s[0])
				sl[1] = append(sl[1], s[1])
			}
			time.Sleep(100 * time.Millisecond)
		}
		sort.Slice(sl[0], func(i, j int) bool {
			return sl[0][i] < sl[0][j]
		})
		sort.Slice(sl[1], func(i, j int) bool {
			return sl[1][i] < sl[1][j]
		})
		mid := STATIC_TIMES >> 1
		if STATIC_TIMES&1 != 0 {
			sum1 += sl[0][mid]
			if staticType == TYPE_C_M {
				sum2 += sl[1][mid]
			}
		} else {
			sum1 += (sl[0][mid] + sl[0][mid-1]) / 2
			if staticType == TYPE_C_M {
				sum2 += (sl[1][mid] + sl[1][mid-1]) / 2
			}
		}
	}
	if staticType == TYPE_CPU {
		result = append(result, strconv.FormatFloat(sum1/3/float64(CPU_COUNT), 'f', -1, 64))
	} else if staticType == TYPE_MEM {
		result = append(result, strconv.FormatFloat(sum1/3, 'f', -1, 64))
	} else {
		result = append(result, strconv.FormatFloat(sum1/3/float64(CPU_COUNT), 'f', -1, 64))
		result = append(result, strconv.FormatFloat(sum1/3, 'f', -1, 64))
	}
	return
}

func captureInfo(pid int, staticType int) ([]float64, error) {
	spid := strconv.Itoa(pid)
	//top -b -n 1 -p $pid  2>&1 | awk -v pid=$pid '{if ($1 == pid)print $9}'
	cmd1 := exec.Command("top", "-b", "-n", "1", "-p", spid)
	//cmd2 := exec.Command("awk","-v","pid="+spid,"{if ($1 == "+spid+")print $9}")
	cmd2 := exec.Command("grep", spid)
	cmd2.Stdin, _ = cmd1.StdoutPipe()
	var cmd3 *exec.Cmd
	if staticType == TYPE_CPU {
		cmd3 = exec.Command("awk", "{print $9}")
	} else if staticType == TYPE_MEM {
		cmd3 = exec.Command("awk", "{print $10}")
	} else if staticType == TYPE_C_M {
		cmd3 = exec.Command("awk", "{print $9,$10}")
	}
	cmd3.Stdin, _ = cmd2.StdoutPipe()
	stdout, _ := cmd3.StdoutPipe()
	//cmd2.Stdout = os.Stdout
	//cmd2.Stderr = os.Stderr

	err := cmd3.Start()
	if err != nil {
		fmt.Println("cmd3 start", err.Error())
		return nil, err
	}
	err = cmd2.Start()
	if err != nil {
		fmt.Println("cmd2 start", err.Error())
		return nil, err
	}
	err = cmd1.Run()
	if err != nil {
		fmt.Println("cmd1 run", err.Error())
		return nil, err
	}
	//err = cmd2.Wait()
	//if err != nil{
	//	fmt.Println("cmd2 wait",err.Error())
	//	return
	//}
	outBytes, err := ioutil.ReadAll(stdout)
	if err != nil || outBytes == nil {
		fmt.Println("reader err", err.Error())
		return nil, err
	}
	result := string(outBytes[:len(outBytes)-1])
	strs := strings.Split(result, " ")
	var r []float64
	for i := range strs {
		//读到的内容包含非法字符串
		for c := range strs[i] {
			if c >= 'a' && c <= 'z' {
				return captureInfo(pid, staticType)
			}
		}
		tmp, err := strconv.ParseFloat(strs[i], 64)
		if err != nil {
			return nil, err
		}
		r = append(r, tmp)
	}
	//如果返回信息为nil
	if r == nil {
		return captureInfo(pid, staticType)
	}
	return r, nil
}
