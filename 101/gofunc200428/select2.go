package main

import (
	"log"
	"time"
)

func main() {

}
func ad() {
	stringChannel := make(chan string, 1)
	go func() {
		time.Sleep(time.Second)
		stringChannel <- "赵卫国"
	}()
	select {
	case <-stringChannel:
		log.Println("stringChannel收到了数据")
	default:
	}
	log.Println("会不会执行，还是继续等待呢")
	time.Sleep(2 * time.Second)
	log.Println(<-stringChannel)
}
