package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	var a []int
	log.Println(a)
	var b bool
	log.Println(b)
	go func() {
		//如果这个线程里面发生指令重排序的话， 就会出现b初始化了，但是a没有初始化，就会panic
		a = make([]int, 3)
		b = true
	}()
	for !b {
		time.Sleep(time.Second)
		runtime.Gosched()
	}
	a[0], a[1], a[2] = 0, 1, 2

}
