package main

import "fmt"

//if 条件判断
func main() {
	age := 19
	if age > 18 {
		fmt.Println("澳门首家线上赌场开业啦")
	} else {
		fmt.Println("该谢暑假作业拉")
	}

	//打印99乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d+%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
