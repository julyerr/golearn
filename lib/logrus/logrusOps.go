package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init(){
	//json 格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//os
	logrus.SetOutput(os.Stdout)
	//level: debug,info,warn,error,fatal,panic
	logrus.SetLevel(logrus.InfoLevel)
}

func main(){
	logrus.WithFields(logrus.Fields{
		"name":"ql",
	}).Info("this is a test")
	logrus.Info("without filed")

	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.WarnLevel)
	logrus.Info("this should not be output")
	logrus.Error("this should be output")

	file,err := os.OpenFile("test.out",os.O_WRONLY|os.O_CREATE,0644)
	if err != nil{
		logrus.Panic(err)
	}
	logrus.SetOutput(file)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Info("this should be output to file ")
}
