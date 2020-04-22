package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
		Zone: "",
	})
	if err != nil {
		fmt.Println("listen eerr", err)
	}

	//不需要建立连接，直接收发数据
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read err", err)
			return
		}
		fmt.Println(data[:n])
		reply := strings.ToUpper(string(data[:n]))
		//发送数据
		conn.WriteToUDP([]byte(reply), addr)
	}
}
