package main

import "fmt"

//浮点数
func main() {

	f1 := 1.123456
	fmt.Printf("%T\n", f1) //默认go语言的小数都是float64类型

	f2 := float32(1.123456)
	fmt.Printf("%T\n", f2) //显示声明float32类型
	//f1 = f2            //float32类型的值不能直接赋值给float64类型的变量
}
