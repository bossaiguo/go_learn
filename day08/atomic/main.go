package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64
var wg sync.WaitGroup

func add() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
