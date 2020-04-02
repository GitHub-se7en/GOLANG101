package main

import (
	"fmt"
	"runtime"
)

func main() {
	//unbuffer()
	buffer()
}
func unbuffer() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int)
	string_chan := make(chan string)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	default:
		fmt.Println("default")
	}

}
func buffer() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}

}
