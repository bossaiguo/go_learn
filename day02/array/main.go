package main

import "fmt"

//数组

//存放元素的容器
//必须指定存放元素的类型和容量
//数组的长度是数组的一部分
func main() {
	var a1 [3]bool //[true, false, true]
	var a2 [4]bool //[true, true, false, false]
	fmt.Printf("a1:%T a2%T\n", a1, a2)

	//初始化方式1
	a3 := [3]bool{true, true, true}
	//初始化方式2
	a4 := [...]int{1, 2, 3, 4, 5}
	//初始方式3
	a5 := [5]int{0: 1, 4: 2}
	a5[3] = 4
	fmt.Println(a3, a4, a5)

	a6 := [...]int{1, 3, 5, 8, 7}
	sum := 0
	for _, v := range a6 {
		sum = sum + v
	}
	fmt.Println(sum)

	//a7 := []int{1, 2, 3}
	//fmt.Println(a7[100]) //会发生越界情况
}
