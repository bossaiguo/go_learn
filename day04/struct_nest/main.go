package main

import "fmt"

//结构体嵌套
type Address struct {
	province string
	city     string
}
type workPlace struct {
	province string
	city     string
}
type Person struct {
	name string
	age  int
	addr Address
}
type Company struct {
	name string
	Address
}

func main() {
	p1 := Person{
		name: "lujing",
		age:  999,
		addr: Address{
			province: "广东",
			city:     "韶关",
		},
	}
	fmt.Println(p1.name, p1.addr)

	c1 := Company{
		name: "越狱",
		Address: Address{
			province: "广东",
			city:     "深圳",
		},
	}
	fmt.Println(c1.Address, c1.city)
}
