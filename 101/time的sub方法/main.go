package main

import (
	"log"
	"time"
)

func main() {
	now := time.Now()
	log.Println(now)
	time.Sleep(time.Second)
	sub := time.Now().Sub(now)
	log.Println(sub)
}
