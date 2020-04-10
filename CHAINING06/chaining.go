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
	//所以这么一来的话，hello就是那个h，真实的处理器其实是return的
	//匿名函数，这个匿名函数包着真正的hello逻辑处理器
	http.HandleFunc("/hello", log(hello))
	server.ListenAndServe()
}
