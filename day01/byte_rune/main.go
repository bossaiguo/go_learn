package main

import "fmt"

//byte和rune类型
//go语言为了处理非ascii码类型的字符，定义了新的rune类型

func main() {
	s := "HELLO沙河"

	for _, c := range s {
		fmt.Printf("%c\n", c) //%c 字符
	}

	//字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2) //把字符串强制转换成了一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3))

	s4 := '红'
	s5 := "红"
	fmt.Printf("s1:%T s5:%T\n", s4, s5)

	//类型转换
	s6 := 10 //int
	var s7 float64
	s7 = float64(s6)
	fmt.Printf("%T\n", s7)
}
