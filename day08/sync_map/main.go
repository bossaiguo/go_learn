package main

import (
	"fmt"
	"strconv"
	"sync"
)

//go 内置的map不是并发安全的
//var m = make(map[string]int, 10)
var k = make([]int, 10)
var m map[string]int

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			//fmt.Printf("k=%v, v=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Printf("%#v\n", m)

	fmt.Println(k)
}
