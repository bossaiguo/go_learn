package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	x int64
}
type result struct {
	*job
	result int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func a(a chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{
			x: x,
		}
		a <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func b(a <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	for {
		job := <-a
		var sum int64 = 0
		n := job.x
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newReuslt := &result{
			job:    job,
			result: sum,
		}
		resultChan <- newReuslt
	}
}

func main() {
	wg.Add(1)
	go a(jobChan)
	for i := 0; i < 24; i++ {
		wg.Add(24)
		go b(jobChan, resultChan)
	}
	for result := range resultChan {
		fmt.Printf("value:%d result:%d\n", result.job.x, result.result)
	}
	wg.Wait()
}
