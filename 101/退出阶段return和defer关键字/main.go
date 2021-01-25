package main

import (
	"fmt"
	"runtime"
)

func f0() int {
	var x = 1
	defer fmt.Println("正常退出", x)
	x++
	//如果是正常退出阶段的话,确实是return之后才进入退出阶段,进入退出阶段才会执行栈里面的退出函数
	return x
}
func f1() int {
	var x, y = 1, 0
	defer fmt.Println("正常退出", x)
	//如果是出现panic的话，return关键字就不会执行了，所以说return和defer的顺序是不一致的哈
	i := x / y
	return i
}
func f2() {
	var x = 1
	defer fmt.Println("正常退出", x)
	//如果一个函数没有返回值的话,执行完下面的语句之后,整个函数就进入了退出阶段
	x++
}

//将在其中的runtime.Goexit函数调用退出完毕之后进入它的退出阶段
//这里的意思是这个函数执行完毕之后,所谓的go里面的函数执行完毕其实是函数退出阶段已经完成,之后的阶段才叫做函数调用退出完毕
func f3() int {
	x := 1
	defer fmt.Println("因Goexit调用而退出:", x)
	x++
	runtime.Goexit()
	return x + x
}

/**
总结就是:
如果是函数正常运行的话,那就是在return语句之后进入函数的退出阶段,
而如果是在return语句之前就发生了panic,那么就会直接进入函数的退出阶段,造成panic的语句之后的所有语句都不会执行
*/
