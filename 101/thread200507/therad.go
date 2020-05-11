package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	asd := "345"
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		asd = "456"
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		log.Println(asd)
		asd = "789"
	}()
	wg.Wait()
}
