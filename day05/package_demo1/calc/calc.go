package calc

import "fmt"

func init() {
	fmt.Println("import 我时，自动执行 calc")
}

func Add(x, y int) int {
	return x + y
}
