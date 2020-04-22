package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile1() {
	fileObj, err := os.Open(`E:\data\go\go\src\go_learn\day05\file_open\main.go`)
	if err != nil {
		fmt.Printf("open file failed:%v", err)
	}
	//记得关闭文件
	defer fileObj.Close()
	//读文件
	var tmp = make([]byte, 128) //指定读的长度
	for {
		n, _ := fileObj.Read(tmp)
		if err != nil {
			fmt.Printf("read fail:%v", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp))
		if n < 128 {
			return
		}
	}
}

func readFromFileByBufio() {
	fileObj, err := os.Open(`E:\data\go\go\src\go_learn\day05\file_open\main.go`)
	if err != nil {
		fmt.Printf("open file failed:%v", err)
	}
	//记得关闭文件
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("read line error:", err)
		}
		fmt.Print(line)
	}
}

func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile(`E:\data\go\go\src\go_learn\day05\file_open\main.go`)
	if err != nil {
		fmt.Println("read file failed err:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}

//打开文件
func main() {
	//readFromFileByBufio()
	readFromFileByIoutil()
}
