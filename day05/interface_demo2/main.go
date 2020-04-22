package main

import "fmt"

//接口的实现
type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步")
}
func (c cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

type chicken struct {
	feet int
}

func (c chicken) move() {
	fmt.Println("鸡冻")
}
func (c chicken) eat() {
	fmt.Println("吃饲料")
}

func main() {
	var a1 animal
	bc := cat{
		name: "淘气",
		feet: 4,
	}
	a1 = bc
	fmt.Printf("%T\n", a1)
	a1.eat("屎")

}
