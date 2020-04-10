package main

import (
	"fmt"
	"reflect"
)

//我升级一下相关的接口

type a interface {
	print()
}
type b struct {
	a
}

func (b *b) print() {
	b.a.print()
}
func test(a a) {
	i := b{a}
	of := reflect.TypeOf(i)
	valueOf := reflect.ValueOf(i)
	name := reflect.Indirect(valueOf).Type().Name()
	fmt.Println(of)
	fmt.Println(valueOf)
	fmt.Println(name)
	i.print()
}
func test2(a a) {
	of := reflect.TypeOf(a)
	valueOf := reflect.ValueOf(a)
	name := reflect.Indirect(valueOf).Type().Name()
	fmt.Println(of)
	fmt.Println(valueOf)
	fmt.Println(name)
	a.print()
}

type c struct {
}

//test一定是需要是接口类型的，为什么外面还需要包一层struct，岂不是多此一举
//
func (c *c) print() {
	fmt.Println("我是c实现的print方法")
}
func main() {
	//test(&c{})
	//记住只有c的指针类型实现了接口，c类型并没有实现
	test2(&c{})
}
