package main

/*
递归函数自己调用自己
递归一定要有一个明确的退出条件， 且越来越接近退出条件
*/

/*
计数n的阶乘
*/
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}
