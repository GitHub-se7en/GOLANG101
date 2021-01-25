package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	time.Sleep(time.Second)
	close(ch)
	log.Println(<-ch)
}
