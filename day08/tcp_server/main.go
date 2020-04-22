package main

import (
	"fmt"
	"net"
)

// tcp server端

func processConn(conn net.Conn) {
	var tmp [128]byte
	n, err := conn.Read(tmp[:])
	if err != nil {
		fmt.Println("read,err", err)
		return
	}
	fmt.Println(string(tmp[:n]))
}

func main() {
	//1.本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("net err")
		return
	}
	//2.等待别人来跟我建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err", err)
			return
		}
		//3.与客户端通信
		go processConn(conn)
	}
}
