package main

import "fmt"

//结构体是值类型

type Person struct {
	name   string
	gender string
}

func f(x Person) {
	x.gender = "女" //修改的是副本的gender
}

func f1(x *Person) {
	(*x).gender = "女" //根据内存地址找到原变量，修改的就是原来的变量
	// x.gender = "不男不女" 语法糖
}

func main() {
	var p Person
	p.name = "lujing"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender)
	f1(&p)
	fmt.Println(p.gender)

	var p2 = new(Person)
	p2.name = "理想"
	p2.gender = "保密"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)

	var p3 = &Person{
		name:   "何去何从",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)

	p4 := &Person{
		"小王子",
		"男",
	}
	fmt.Printf("%#v\n", p4)
}
