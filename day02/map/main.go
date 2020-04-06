package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        //还没初始化（没有在内存种开辟空间）
	m1 = make(map[string]int, 10) //要估算好该map容量， 避免在程序运行期间再动态扩容
	m1["a"] = 1
	m1["b"] = 2

	fmt.Println(m1["娜扎"]) //拿到对应类型的0值

	v, ok := m1["娜扎"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(v)
	}

	//删除
	delete(m1, "a")
}
