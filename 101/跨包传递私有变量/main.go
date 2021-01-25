package main

import (
	"GOLANG101/101/跨包传递私有变量/inner"
	"log"
)

func main() {
	a := inner.A()
	log.Println(a)
	/**
	直接点是无法访问到小写的属性的,
	但是这样的话,其实也是实现了一种效果,那就是外界只能看,不能访问,不能修改,
	只能是提供工厂函数来
	*/
	//log.Println(a.)
	type name struct {
	}
	inner.B(name{})

	type name2 func()
	inner.C(name2(s))
}
func s() {

}
