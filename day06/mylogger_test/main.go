package main

import "go_learn/day06/mylogger"

func main() {
	mylogger.NewLog("debug").Debug("这是一条debug日志")

}
