package main

import "fmt"

//make()函数创造切片
/*
	1. 要判断一个切片是否是空， 要用 len(s) === 0 来判断， 不应该使用 s == nil 来判断
*/
func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1-%v len(s1)-%d cap(s1)-%d", s1, len(s1), cap(s1))
}
