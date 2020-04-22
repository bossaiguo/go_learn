package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())

	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	//time.Unix()
	ret := time.Unix(1586596049, 0)
	fmt.Println(ret)

	//now+1小时
	fmt.Println(now.Add(24 * time.Hour))

	//定时器
	/*timer := time.Tick(time.Second)
	for t := range timer {
		fmt.Println(t)
	}*/

	//格式话时间 把语言中时间对象 转会成字符串类型的时间
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("06/01/02 03:04:05"))

	//安卓对应的格式解析字符串类型的时间
	timeObjs, _ := time.Parse("2006-01-02", "2019-01-12")
	fmt.Println(timeObjs)

	//sleep
	n := 5 //int
	time.Sleep(time.Duration(n) * time.Second)
	time.Sleep(5 * time.Second)
}
