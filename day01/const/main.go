package main

//iota 在const关键字出现时将被重置为0，const中每新增一行常量声明将使iota计数一次

//插队
const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4
)

//多行常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 //d1:1 d2:2
	d3, d4 = iota + 1, iota + 2 //d3:2 d4:3
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
