package _func

import (
	"fmt"
	"log"
)

func init() {
	fmt.Println("我是init方法，会执行")
}

func Function() {
	var aa func()
	aa = ss
	//log.Println(aa)
	aa()
	var bb = func() { log.Println("我是bb函数") }
	bb()
}
func ss() {
	log.Println("我是ss函数")
}
