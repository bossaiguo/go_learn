package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "lujing", "请输入名字")
	flag.Parse()
	fmt.Println(*name)
}
