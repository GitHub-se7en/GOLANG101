package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO")
}
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("调用的处理器函数的名字是", name)
		h(w, r)
	}
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandlerFunc("").
		server.ListenAndServe()
}
