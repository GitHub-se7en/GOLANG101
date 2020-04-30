package main

import (
	"log"
	"time"
)

func a() string {
	go func() {
		log.Println("我是支线在执行任务")
	}()
	return "我是主线"
}
func main() {
	//按理说主线结束之后，支线任务就不应该继续执行了啊
	log.Println(a())
	time.Sleep(time.Second)
}
