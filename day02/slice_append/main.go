package main

import "fmt"

//append() 为切片追加元素

func main() {
	s1 := []int{1, 2, 3}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	//s1[3] = 4 //错误写法， 会导致编译错误：索引越界

	//调用append函数必须用原来的切片变量接收返回值
	s1 = append(s1, 4)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s2 := []int{5, 6, 7, 8, 9, 10, 11, 12, 13}
	s1 = append(s1, s2...) // ...表示拆开
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}
