package main

import (
	"bufio"
	"fmt"
	"os"
)

func useScan() {
	var s string
	fmt.Print("请输入内容：")
	fmt.Scanln(&s)
	fmt.Printf("你输入的那内容时：%s\n", s)
}

func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的那内容是：%s\n", s)
}

func main() {
	//useScan()
	useBufio()
}
