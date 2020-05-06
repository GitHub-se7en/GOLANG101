package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("我是支线在执行任务")
		time.Sleep(2 * time.Second)
		return
		select {
		case <-time.After(time.Second):
			log.Println("时间到了")
			return
		}
	}()
	wg.Wait()
}
