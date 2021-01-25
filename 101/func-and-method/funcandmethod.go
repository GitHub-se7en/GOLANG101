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
	/**
	在这里面有两个的例子，一个是如果是interface的话，那就直接点类型 ，就能转型成功，
	第二个如果是函数类型的话，没有点，后面直接追加括号就可以实现转型
	*/
	handlerFunc := HandlerFunc(name)
	handler := Handler(handlerFunc)
	handler.(HandlerFunc).ServeHTTP()
	handlerFunc.ServeHTTP()
	s := string("saddfd")
	student := student(student{name: "zhaoweiguo"})
}
func name(name string) {
	fmt.Println("姓名", name)
}
