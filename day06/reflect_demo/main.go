package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type name:%v type kind:%v\n", v.Name(), v.Kind())
}

func main() {
	/*str := `{"name":"卢晶","age":9000}`
	var p Person
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age)*/

	var c = Person{
		Name: "lujing",
		Age:  9000,
	}
	reflectType(c)
}
