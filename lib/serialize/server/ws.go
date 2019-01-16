package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/flatbuffers/go"
	"github.com/json-iterator/go"
	"go.megvii-inc.com/securitycore/face/cmd/capture-benchmark/Capture/Message"
	"go.megvii-inc.com/securitycore/face/cmd/capture-benchmark/capM"
	"golearn/lib/serialize/capM"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	flagPath = flag.String("path","","image path")
	flagRate = flag.Int("fps", 5, "send image fps")
	flagPort = flag.String("port", "9003", "listen port")
	flagType = flag.Int("type",0,"type of message")
)

type captureMessage struct {
	CropImage  []byte
	FaceImage  []byte
	Timestamp      int64  `json:"timestamp"`
	Track          int    `json:"track"`
	Type           string `json:"type"`
	SeqeunceNumber int64  `json:"seq_num"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfo,err := os.Stat(*flagPath)
	if err != nil{
		panic(err)
	}
	filenames := []string{}
	if fileInfo.IsDir(){
		fs,_:= ioutil.ReadDir(*flagPath)
		for i:=0;i<len(fs);i++{
			filenames = append(filenames,*flagPath+string(os.PathSeparator)+fs[i].Name())
		}
	}else{
		filenames = append(filenames,*flagPath)
	}
	_,message,err := conn.ReadMessage()
	if err != nil{
		return
	}
	//fmt.Println("server receive type",string(message))
	mType,err := strconv.Atoi(string(message))
	if err != nil{
		return
	}
	for{
		for i:= 0 ;i< len(filenames);i++{
			t := time.Now()
			data,err := ioutil.ReadFile(filenames[i])
			if err != nil{
				panic(err)
			}
			b,err := marshalData(mType,data)
			if err != nil{
				fmt.Println("server marshal failed","type",mType,"err",err)
				return
			}
			err = conn.WriteMessage(websocket.BinaryMessage,b)
			if err != nil{
				//fmt.Println("server send failed",err)
				return
			}
			if lag := time.Second/time.Duration(*flagRate) - time.Now().Sub(t); lag > time.Duration(0) {
				time.Sleep(lag)
			} else if lag < time.Duration(0) {
				//logrus.Warnf("Send too slow: %v", lag)
			}
		}
	}
}

func marshalData(mType int,data []byte) ([]byte,error){
	switch {
	case mType == 0:
		cap := &captureMessage{
			Type:      "capture",
			Timestamp: int64(time.Now().Second()),
			CropImage: data,
			FaceImage: data,
		}
		b,err := json.Marshal(cap)
		return b,err
	case mType == 1:
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		cap := &captureMessage{
			Type:      "capture",
			Timestamp: int64(time.Now().Second()),
			CropImage: data,
			FaceImage: data,
		}
		b,err := json.Marshal(cap)
		return b,err
	case mType == 2:
		//pb string must contains valid utf-8
		captureMessage := &capture_message.CaptureMessage{
			CropImage: data,
			FaceImage: data,
			TimeStamp: int64(time.Now().Second()),
			Type: "capture",
		}
		b,err := proto.Marshal(captureMessage)
		return b,err
	case mType == 3:
		builder := flatbuffers.NewBuilder(0)
		capType := builder.CreateString("capture")
		cropImage := builder.CreateString(string(data))
		faceImage := cropImage
		//size := len(data)
		//Message.CaptureMessageStartCapImageVector(builder,size)
		//for i:= size -1 ;i>=0;i--{
		//	builder.PrependByte(data[i])
		//}
		//cropImage := builder.EndVector(size)
		//
		//Message.CaptureMessageStartCapImageVector(builder,size)
		//for i:= size -1 ;i>=0;i--{
		//	builder.PrependByte(data[i])
		//}
		//faceImage := builder.EndVector(size)
		Message.CaptureMessageStart(builder)
		times := int64(time.Now().Second())
		Message.CaptureMessageAddTimestamp(builder,times)
		Message.CaptureMessageAddCapImage(builder,cropImage)
		Message.CaptureMessageAddFaceImage(builder,faceImage)
		Message.CaptureMessageAddType(builder,capType)
		end := Message.CaptureMessageEnd(builder)
		builder.Finish(end)

		buf := builder.FinishedBytes()

		//md := md5.New()
		//fmt.Println("-----before","capType:","capture","times:",times,"md5:",fmt.Sprintf("%x",md.Sum(data)[:15]))
		//
		//capMessage := Message.GetRootAsCaptureMessage(buf,0)
		//capType1 := string(capMessage.Type())
		//times1 := capMessage.Timestamp()
		//size1 := capMessage.CapImageLength()
		//newData := make([]byte,size1)
		//for i:=0;i<size1;i++{
		//	newData[i] = byte(capMessage.CapImage(i))
		//}
		//md.Reset()
		//fmt.Println("-----after","capType:",capType1,"times:",times1,"md5:",fmt.Sprintf("%x",md.Sum(newData)[:15]))
		return buf,nil
	default:
	}
	return nil,errors.New("not supported message type")
}

func main() {
	flag.Parse()
	if *flagPath == ""{
		fmt.Println("image path")
		return
	}

	//go func() {
	//	for {
	//		time.Sleep(1 * time.Second)
	//		fpsMu.Lock()
	//		logrus.Infoln("FPS:", fps)
	//		fps = 0
	//		fpsMu.Unlock()
	//	}
	//}()

	http.HandleFunc("/test", handler)
	panic(http.ListenAndServe("0.0.0.0:"+*flagPort, nil))
}

