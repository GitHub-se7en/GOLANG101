package main

import "log"

func main() {
	b := '\r'
	if b != '\r' || b != '\n' {
		log.Println("asdfafdadfasf")
	}
	//关键是判断条件等于换成了不等于，所以出问题了，转不过来了
	//解决方法就是a等于1不就等于a不等于2吗，出结论
	a := 1
	if a != 1 || a != 2 {
		log.Println("asdfafdadfasf")
	}
}
