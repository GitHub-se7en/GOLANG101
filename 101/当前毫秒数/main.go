package main

import (
	"log"
	"time"
)

func main() {
	//毫秒数
	i := time.Now().UnixNano() / 1e6
	//纳秒数
	j := time.Now().UnixNano()
	log.Println(i)
	log.Println(j)
}
