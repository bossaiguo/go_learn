package main

import "fmt"

/*
	1. &取地址
	2. *根据地址取值
*/
func main() {
	// 1. &取地址
	n := 18
	p := &n
	fmt.Println(&n)
	fmt.Printf("%T\n", p)

	//2. *根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	//new函数申请一个内存地址
	var a1 *int = new(int)
	fmt.Println(a1)
	fmt.Println(*a1)
	*a1 = 100
	fmt.Println(*a1)

	/*
		make也是用于内存分配的，区别于new， 它只用于slice、map及channel的内存创建，而且他返回的类型就是
		这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
	*/
	/*
		1.make和new都是用来申请内存的
		2.new很少用，一般用来给基本数据类型申请内存， string\int... 返回的是对应类型的指针
		3.make是用来给slice、map、channel申请内存的，make函数返回的是对应的这三个类型本身
	*/
}
