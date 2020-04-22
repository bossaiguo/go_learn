package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseInt("11", 10, 64))
	fmt.Println(strconv.Atoi("22"))
	fmt.Println(strconv.Itoa(33))

	fmt.Println(strconv.ParseBool("true"))
}
