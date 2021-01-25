package main

import "log"

func main() {
	a(b(c(total)))   //执行这个函数的结果就是返回了一个不断嵌套的函数，因为这个函数里面没有执行语句，只有return，所以没有输出
	a(b(c(total)))() //这个就是执行了匿名函数了，其实它可以返回一个函数的，然后在执行这个函数
	//这样就模拟出了嵌套函数的形式
}
func a(reqA func()) func() {
	return func() {
		log.Println("这个是a函数里面的日志")
		reqA()
	}
}
func b(reqB func()) func() {
	return func() {
		log.Println("这个是b函数里面的日志")
		reqB()
	}
}
func c(reqC func()) func() {
	return func() {
		log.Println("这个是c函数里面的日志")
		reqC()
	}
}

//目标是创建出一个无限套用的函数，层层嵌套的函数
func total() {
	log.Println("这个是total函数里面的日志")
}
