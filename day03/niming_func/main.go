package main

import "fmt"

/*
匿名函数
函数内部无法声明带名字的函数
*/

func main() {
	var f1 = func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)
}
