package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("dial err", err)
		return
	}
	defer socket.Close()
	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入内容：")
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))
		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("reply err", err)
		}
		fmt.Println("收到回复信息", string(reply[:n]))
	}
}
