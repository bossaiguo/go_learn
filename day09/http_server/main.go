package main

import "net/http"

func f1(w http.ResponseWriter, r *http.Request) {
	str := "hello dali"
	w.Write([]byte(str))
}
func main() {
	http.HandleFunc("/hello", f1)
	http.ListenAndServe("127.0.0.1:9000", nil)
}
