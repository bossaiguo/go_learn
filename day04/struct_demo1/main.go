package main

import "fmt"

//结构体

type Person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var lujing Person
	lujing.name = "lujing"
	lujing.age = 9000
	lujing.gender = "男"
	lujing.hobby = []string{"历史", "哲学"}
	fmt.Println(lujing)

	//匿名结构体：多用于临时场景
	var s struct {
		x string
		y int
	}
	s.x = "嘿嘿嘿"
	s.y = 100
	fmt.Println("type:%T value:%v\n", s, s)
}
