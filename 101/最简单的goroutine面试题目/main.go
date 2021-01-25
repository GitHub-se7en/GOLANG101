package main

import (
	"log"
	"math/rand"
	"sync"
)

func main() {
	//开启两个协程,用channel进行通信
	strings := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			strings <- rand.Intn(5)
		}
		close(strings)
	}()
	go func() {
		defer wg.Done()
		for value := range strings {
			log.Println(value)
		}
	}()
	wg.Wait()
}
