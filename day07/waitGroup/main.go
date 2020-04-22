package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	1.并发：同一时间段内执行多个任务
	2.并行：同一个时刻执行多个任务
*/
/*
	M:N 把m个goroutine 分配给n个操作系统线程去执行
*/
func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Printf("r1:%v r2:%v\n", r1, r2)
	}
}

func f1(i int) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Second * time.Duration(rand.Intn(20)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	wg.Wait()
}
