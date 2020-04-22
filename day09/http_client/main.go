package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	r, err := http.Get("http://127.0.0.1:9000/hello")
	if err != nil {
		fmt.Printf("get err:%v\n", err)
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read body %v\n", err)
		return
	}
	println(string(b))

}
