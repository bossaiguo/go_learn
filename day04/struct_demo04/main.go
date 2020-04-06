package main

import "fmt"

//方法
type dog struct {
	name string
	age  int
}

//构造函数
func newDog(name string, age int) dog {
	return dog{
		name: name,
		age:  age,
	}
}

/*
方法是作用域特定类型的函数
接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
*/
//值接收者
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪", d.name)
}

/*
	什么时候应该使用指针类型接收者？
	1. 需要修改接收者的值
	2. 接收者是拷贝代价比较大的大对象
	3. 保持一致性， 如果有某个方法使用了指针接收者，那么其他方法也应该使用指针接收者
*/
func (d *dog) guonian() {
	d.age++
}

func main() {
	d1 := newDog("lujing", 25)
	d1.wang()
	d1.guonian()
	fmt.Println(d1.age)
}
