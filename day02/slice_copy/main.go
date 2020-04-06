package main

import (
	"fmt"
	"sort"
)

func main() {
	a1 := []int{1, 2, 3}
	a2 := a1
	a3 := make([]int, 3, 3)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	//将a1种的索引为1的元素删除
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)

	a4 := make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a4 = append(a4, i)
	}
	fmt.Println(a4)

	a5 := [...]int{3, 7, 8, 9, 1}
	sort.Ints(a5[:]) //对切片进行排序
	fmt.Println(a5)
}
