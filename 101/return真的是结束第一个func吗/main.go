package main

import (
	"fmt"
	"time"
)

func main() {
	w := make(chan struct{}, 3)
	//s := make(chan struct{},3)
	var s chan struct{}
	w <- struct{}{}
	worker(w, s)
	s = make(chan struct{}, 3)
	s <- struct{}{}
	time.Sleep(time.Second)
}

/**
我一开始以为的是如果执行了一个select case的话,这个结构会不会继续监听别的case,但是现实是不是的,这个case没有执行完的话,是不会去判断另外一个case的
这里面其实就是假设了如果是启动了一个任务的话,是没办法中途停止的,除非是这个任务结束之后,执行下一个任务之前会判断一次
*/
func worker(w <-chan struct{}, s <-chan struct{}) {
	go func() {
		defer fmt.Println("worker exit")
		// Using stop channel explicit exit
		for {
			select {
			case <-s:
				fmt.Println("Recv stop signal")
				return
			case <-w:
				fmt.Println("Working .")
			}
		}
	}()
	return
}
