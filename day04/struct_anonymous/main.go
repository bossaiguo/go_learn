package main

import "fmt"

//匿名字段
//不常用
type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		"lujing",
		9000,
	}
	fmt.Println(p1.string)
}
