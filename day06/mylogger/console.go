package mylogger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint16

const (
	DEBUG LogLevel = iota
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger struct {
	level LogLevel
}

func parseLogLevel(s string) LogLevel {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG
	case "trace":
		return TRACE
	case "info":
		return INFO
	case "warning":
		return WARNING
	case "fatal":
		return FATAL
	default:
		return DEBUG
	}
}

// 构造函数
func NewLog(level string) Logger {
	levelLog := parseLogLevel(level)
	return Logger{
		level: levelLog,
	}
}

func log(lv LogLevel, msg string) {
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	nowTime := now.Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [DEBUG] [%s:%s:%d] %s\n", nowTime, fileName, funcName, lineNo, msg)
}

func (l Logger) Debug(msg string) {

}

func (l Logger) Info(msg string) {
	fmt.Println(msg)
}

func getInfo(n int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("runtime.Caller() failed\n")
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}
