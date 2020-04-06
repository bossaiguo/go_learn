package main

import "unicode"

//1. 判断字符串中汉字的数量
//难点是判断一个字符是否是汉字
func main() {
	s1 := "Hello World 深圳"
	//1. 一次拿到字符串中的字符
	//2. 判断当前这个字符是否是汉字
	//3. 把汉字出现的次数累加得到最终结果
	var count int
	for _, c := range s1 {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	println(count)
}
