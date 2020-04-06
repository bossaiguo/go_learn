package main

import "fmt"

func main() {
	//var s1 [2]int
	//切片的定义
	/*
		1.切片是指向了一个底层数组
		2.切片的长度就是它元素的个数
		3.切片的容量是底层数组从切片第一个元素导最后一个元素的数量
		4.切片是引用类型，都指向了一个底层数组
	*/
	var s2 []int    //定义一个存放int类型的切片
	var s3 []string //定义一个存放string类型元素的切片

	//fmt.Println(s1 == nil)
	fmt.Println(s2 == nil) //没有开辟内存空间
	fmt.Println(s3 == nil)

	s2 = []int{1, 2, 3}
	//长度和容量
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	//2.有数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6}
	s4 := a1[0:4] //基于一个数组切割， 左闭右开
	fmt.Println(s4)

	s5 := a1[:4]
	s6 := a1[3:]
	s7 := a1[:]
	fmt.Println(s5, s6, s7)
}
