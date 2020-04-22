package main

import (
	"log"
	"os"
	"time"
)

func main() {
	fileObj, _ := os.OpenFile("./test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(fileObj)
	for {
		log.Println("这是一条测试日志")
		time.Sleep(3 * time.Second)
	}
}
