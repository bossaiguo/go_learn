package main

import (
	"fmt"
	"strings"
)

func main() {
	//字符串相关操作
	s3 := "hello world"

	//字符串相关操作
	fmt.Println(len(s3))

	//字符串拼接
	name := "理想"
	word := "dsb"
	s4 := name + word
	s5 := fmt.Sprintf("%s%s", name, word)
	fmt.Println(s4)
	fmt.Println(s5)

	//字符串分割
	ret := strings.Split(s3, " ")
	fmt.Println(ret)

	//包含
	fmt.Println(strings.Contains(s4, "理想"))

	//前缀
	fmt.Println(strings.HasPrefix(s4, "理想"))

	//后缀
	fmt.Println(strings.HasSuffix(s4, "理想"))

	//位置
	fmt.Println(strings.Index(s4, "d"))
	fmt.Println(strings.LastIndex(s4, "d"))

	//拼接
	fmt.Println(strings.Join(ret, "+"))
}
