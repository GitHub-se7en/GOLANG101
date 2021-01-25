package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	//谁说的一定是要有panic才会触发defer关键字的，那个王八蛋说的，贱人
	//panic("触发异常")
}
