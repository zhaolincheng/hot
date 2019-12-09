package util

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger // 记录所有日志
	Info    *log.Logger // 记录重要信息
	Warning *log.Logger // 记录注意信息
	Error   *log.Logger // 记录错误信息
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	Trace = log.New(ioutil.Discard, "Trace: ", log.Ldate|log.Ltime|log.Llongfile)
	Info = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Llongfile)
	Warning = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Llongfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "Error: ", log.Ldate|log.Ltime|log.Llongfile)
}
