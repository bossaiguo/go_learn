package main

import "fmt"

/*
断言
*/
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println(str)
	}
}

func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串", t)
	case int:
		fmt.Println("是一个int", t)

	}
}
func main() {
	m1 := make(map[string]int, 10)
	m2 := make([]int, 10, 10)
	fmt.Println(m1)
	fmt.Println(m2)

	//assign("asd")
	assign2("好好学习")
}
