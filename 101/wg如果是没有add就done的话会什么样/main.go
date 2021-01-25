package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	//wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("我是最标准的协程")
	}()
	//wg.Wait()
	time.Sleep(time.Second)
	log.Println("等待所有的协程完成之后,在完成这个工作")
}
