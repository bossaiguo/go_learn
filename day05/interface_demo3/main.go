package main

import "fmt"

/*
使用值接收者和指针接收者的区别
1.使用值接收者实现接口，结构体类型和结构指针类型的变量都能存。
2.指针接收者实现接口只能存结构体指针类型的变量

接口和指针的关系
1.多个类型可以实现同一个接口
2.一个类型可以实现多个接口

注意
只有当有两个或两个以上的具体类型必须以相同方式进行处理时才需要定义接口，
不要为了接口而写接口，这样只会增加不必要的抽象，导致不必要的运行时消耗
*/

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c *cat) move() {
	fmt.Println("走猫步")
}
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

func main() {
	//var a1 animal
	c1 := cat{"tom", 4}
	c2 := &cat{"假老练", 5}

	c1.move()
	c2.eat("mouse")
}
