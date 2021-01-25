package main

import (
	"log"
	"sync/atomic"
)

func main() {
	var count uint32
	val := atomic.LoadUint32(&count)
	log.Println(val)
}
