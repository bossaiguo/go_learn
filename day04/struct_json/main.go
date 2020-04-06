package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age" db:"age" ini:"age"`
}

func main() {
	p1 := Person{
		Name: "卢晶",
		Age:  9000,
	}
	//序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Printf("%#v\n", string(b))
	//反序列化
	str := `{"name":"理想","age":18}`
	var p2 Person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)
}
