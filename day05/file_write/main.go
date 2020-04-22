package main

import (
	"fmt"
	"os"
)

func main() {
	os.OpenFile(`E:\data\go\go\src\go_learn\day05\file_write\test.txt`, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
	}
}
