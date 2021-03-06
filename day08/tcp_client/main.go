package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//1.与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial err", err)
	}

	//2.发送数据
	var msg string
	if len(os.Args) < 2 {
		msg = "hello world"
	} else {
		msg = os.Args[1]
	}
	conn.Write([]byte(msg))
	conn.Close()
}
