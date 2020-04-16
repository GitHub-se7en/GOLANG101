package main

import "log"

type a interface {
	printZhaoWeiGuo()
}
type b interface {
	printZhaoWeiGuo()
	printXiaoming()
}
type f struct {
}

func (f *f) printXiaoming() {
	log.Println("小明")
}
func (f *f) printZhaoWeiGuo() {
	log.Println("赵卫国")
}

//所有的数据都在listenTcp的时候确定了，后面所有的接口都是这个TCP实现的具体类
func main() {
	a(b(&f{})).printZhaoWeiGuo()
}
