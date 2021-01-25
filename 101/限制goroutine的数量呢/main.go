package main

import (
	"log"
	"sync"
	"time"
)

/**
不好限制数量,基本上都是使用go关键字直接启动协程的,但是怎么限制这个协程的数量就是一个很大的问题啊
*/
func main() {
	//wg := sync.WaitGroup{}
	//f1() //第一个版本
	f2() //第二个版本
}
func f1() {
	var wg1 sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg1.Add(1)
		go func(i int) {
			log.Printf("我是第%v个执行的线程", i)
			wg1.Done()
		}(i)
	}
	wg1.Wait()
}
func f2() {
	//怎么限制goroutine的数量呢?channel是用来通信的,两个协程之间进行通信
	var wg1 sync.WaitGroup
	communication := make(chan struct{}, 5)
	for i := 1; i <= 5; i++ {
		//初始化协程池
		communication <- struct{}{}
	}
	for i := 1; i <= 10; i++ {
		wg1.Add(1)
		<-communication
		go func(i int) {
			log.Printf("我是第%v个执行的线程", i)
			//模拟协程负载
			time.Sleep(time.Second)
			communication <- struct{}{}
			wg1.Done()
		}(i)
	}
	wg1.Wait()
}

/**
我之前怎么没看出来,这他娘的是一对多的关系啊,一个线程和多个线程之间进行通信啊,但是我就是始终想不出来是怎么和限制协程数量关系起来的啊
*/
