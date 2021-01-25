package main

import (
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	for i := 1; i <= 10; i++ {
		log.Println(rand.Intn(100))
	}
}
