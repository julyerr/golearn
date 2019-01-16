package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/json-iterator/go"
	"golearn/lib/serialize/Capture/Message"
	"golearn/lib/serialize/capM"
)

const (
	CPU_COUNT = 56
)

const (
	TYPE_CPU = iota
	TYPE_MEM
	TYPE_C_M
)

var (
	flagConcurrency = flag.Int("con", 100, "concurrency to create video stream for each websocket server")
	flagRate        = flag.Int("fps", 5, "rate of every stream")
	flagFilename    = flag.String("name", "result.txt", "save filename")
	flagCputimes    = flag.Int("times", 3, "pid")
	flagStaticType  = flag.Int("type", 0, "static type(0 -> cpu,1 -> mem,2 -> cpu,mem)")
	mType           int
	dialer          *websocket.Dialer
	u               = url.URL{Scheme: "ws", Host: "10.199.0.249:" + strconv.Itoa(9003), Path: "/test"}
	pid             int
)

type captureMessage struct {
	CropImage      []byte
	FaceImage      []byte
	Timestamp      int64  `json:"timestamp"`
	Track          int    `json:"track"`
	Type           string `json:"type"`
	SeqeunceNumber int64  `json:"seq_num"`
}

func process(serverNumber int, exit chan int, staticType int) ([]string, error) {
	stop := make(chan int)
	for j := 0; j < serverNumber; j++ {
		go func() {
			conn, _, err := dialer.Dial(u.String(), nil)
			if err != nil {
				return
			}
			err = conn.WriteMessage(websocket.BinaryMessage, []byte(strconv.Itoa(mType)))
			if err != nil {
				return
			}
			for {
				select {
				case <-exit:
					return
				case <-stop:
					return
				default:
				}

				select {
				case <-exit:
					return
				case <-stop:
					return
				default:
					_, message, err := conn.ReadMessage()
					if err != nil {
						fmt.Println("client read:", err)
						return
					}
					err = unmarshal(mType, message)
					if err != nil {
						fmt.Println("unmarshal failed", err, "type", mType)
						return
					}
				}
			}
		}()
	}
	time.Sleep(1 * time.Second)
	s, e := staticsInfo(staticType)
	if e != nil {
		return nil, e
	}
	close(stop)
	return s, nil
}

func unmarshal(mType int, data []byte) error {
	switch mType {
	case 0:
		var cap captureMessage
		err := json.Unmarshal(data, &cap)
		return err
	case 1:
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		var cap captureMessage
		err := json.Unmarshal(data, &cap)
		return err
	case 2:
		var captureMessage capture_message.CaptureMessage
		err := proto.Unmarshal(data, &captureMessage)
		//f,err := os.OpenFile("test.jpg",os.O_WRONLY,0644)
		//defer f.Close()
		//if err != nil{
		//	fmt.Println("open file failed")
		//	return err
		//}
		//f.Write(captureMessage.FaceImage)
		//fmt.Println(captureMessage.GetType())
		return err
	case 3:
		capMessage := Message.GetRootAsCaptureMessage(data, 0)
		//得到的bytes
		capMessage.CapImage()
		capMessage.FaceImage()
		//md := md5.New()
		//fmt.Println("-----after","capType:",string(capMessage.Type()),"times:",capMessage.Timestamp(),"md5:cropImage",fmt.Sprintf("%x",md.Sum(cropImage)[:15]))
		//capType := string(capMessage.Type())
		//timeStamp := capMessage.Timestamp()
		//size1 := capMessage.CapImageLength()
		//newData := make([]byte, size1)
		//for i := 0; i < size1; i++ {
		//	newData[i] = byte(capMessage.CapImage(i))
		//}
		//size2 := capMessage.FaceImageLength()
		//newData2 := make([]byte, size2)
		//for i := 0; i < size2; i++ {
		//	newData2[i] = byte(capMessage.FaceImage(i))
		//}
		////fmt.Println(string(capMessage.Type()))
		//return nil
		return nil
	default:
	}
	return errors.New("not supported message type")
}

func staticsInfo(staticType int) (result []string, err error) {
	var sum1, sum2 float64
	for i := 0; i < 3; i++ {
		sl := make([][]float64, 2)
		for i := 0; i < *flagCputimes; i++ {
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
		mid := *flagCputimes >> 1
		if *flagCputimes&1 != 0 {
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
	if err != nil {
		fmt.Println("reader err", err.Error())
		return nil, err
	}
	result := string(outBytes[:len(outBytes)-1])
	strs := strings.Split(result, " ")
	var r []float64
	for i := range strs {
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

func main() {
	//参数校验
	flag.Parse()
	if *flagStaticType < 0 || *flagStaticType > 2 {
		fmt.Println("static type must be mem or cpu")
		return
	}
	pid = os.Getpid()

	file1, err := os.OpenFile(*flagFilename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	fmt.Fprintf(file1, "concurrent:%d , rate:%d\n", *flagConcurrency, *flagRate)
	defer file1.Close()
	if err != nil {
		panic(err)
	}
	var file2 *os.File
	if *flagStaticType == TYPE_C_M {
		file2, err = os.OpenFile(*flagFilename+"_ano", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		fmt.Fprintf(file1, "concurrent:%d , rate:%d\n", *flagConcurrency, *flagRate)
		defer file2.Close()
		if err != nil {
			panic(err)
		}
	}
	var types []string
	types = append(types, "json")
	types = append(types, "json-ter")
	types = append(types, "pb")
	types = append(types, "fb")

	//处理中断
	c := make(chan os.Signal, 1)
	done := make(chan int, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		close(done)
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}()

	uplimit := *flagConcurrency / 10
	if uplimit <= 0 {
		uplimit = 1
	}
	for i := 0; i < 4; i++ {
		start := 1
		mType = i
		file1.WriteString("testing type " + types[mType] + " \n")
		for {
			s, e := process(start, done, *flagStaticType)
			if e != nil {
				panic(e)
			}
			str := strconv.Itoa(start**flagRate) + "\t" + s[0]
			fmt.Println(str)
			file1.WriteString(str + "\n")
			if *flagStaticType == TYPE_C_M {
				str := strconv.Itoa(start**flagRate) + "\t" + s[1]
				fmt.Println(str)
				file2.WriteString(str + "\n")
			}
			if start > uplimit {
				start += uplimit
			} else {
				start += start
			}
			if start > *flagConcurrency {
				fmt.Println("finished ,outbreak")
				break
			}
		}
		time.Sleep(1 * time.Second)
	}
	return
}
