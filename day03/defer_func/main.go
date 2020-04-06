package main

import "fmt"

/*
在go语言中，return语句底层并不是原子操作，他分为给返回值赋值和RET指令
两步， 而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
*/
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("a")
	defer fmt.Println("b")
	fmt.Println("end")
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++ //改变的是函数中的副本
	}(x)
	return 5
}

func main() {
	deferDemo()
}
