package main

import "fmt"

type Handler interface {
	ServeHTTP()
}

type HandlerFunc func(name string)

type student struct {
	name string
}

func (h HandlerFunc) ServeHTTP() {
	h("赵卫国")
	fmt.Println("我是h实现的ServeHTTP方法")
}
func main() {
	handlerFunc := HandlerFunc(name)
	handler := Handler(handlerFunc)
	handlerFunc.ServeHTTP()
	s := string("saddfd")
	student := student(student{name: "zhaoweiguo"})
}
func name(name string) {
	fmt.Println("姓名", name)
}
