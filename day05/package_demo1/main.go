package main

import (
	"fmt"
	"go_learn/day05/package_demo1/calc"
)

/*
包中的标识符（变量名/函数名/结构体/）如果首字母是小写的，表示私有
*/
func main() {
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}
