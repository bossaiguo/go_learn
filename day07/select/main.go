package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 1000; i++ {
		select {
		case x := <-ch:
			fmt.Printf("取出%v\n", x)
			time.Sleep(time.Millisecond * 2000)
		case ch <- i:
			fmt.Printf("插入%v\n", i)
		}
	}
}
