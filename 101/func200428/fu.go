package main

import "log"

type Func func(i int)

func add(int2 int) {
	log.Println("函数的参数是----", int2)
}

type hasFunc struct {
	hasFunc Func
}

func main() {
	//也就是说和之前的TCP一样，在初始化的时候就把具体的类型固定了
	h := hasFunc{hasFunc: add}
	h.hasFunc(5)
	i := Func(add)
	i(6)
}
