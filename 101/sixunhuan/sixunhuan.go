package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("是不是死循环")
	}
}
