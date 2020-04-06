package main

import "fmt"

//函数类型
func f1() {
	fmt.Println("Hello World")
}

//函数作为参数的类型
func f3(x func() int) {
	x()
}

//函数还可以作为返回值
func f5(x func() int) func(int, int) int {
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}
func main() {
	a := f1
	fmt.Printf("%T\n", a)
}
