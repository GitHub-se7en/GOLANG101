package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	// Set a duration.
	duration := 150 * time.Millisecond
	// Create a context that is both manually cancellable and will signal
	// cancel at the specified duration.
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// Create a channel to receive a signal that work is done.
	ch := make(chan string)

	// Ask the goroutine to do some work for us.
	go func() {
		//先启动一个监听器，然后在执行代码逻辑，如果代码逻辑出现问题，比如死循环之类的，监听器到期之后会自动结束这个函数
		select {
		case <-ctx.Done():
			log.Println("上方调用取消了这个函数")
			return
		}
		// Simulate work.
		time.Sleep(250 * time.Millisecond)
		// Report the work is done.
		ch <- "123"
		//我需要你做一件事，我才能给你一些东西
		//我需要你给我这些东西，我才做这件事
		//我需要你给我这些东西，但是我不会做这件事
	}()

	// Wait for the work to finish. If it takes too long, move on.
	select {
	case d := <-ch:
		fmt.Println("work complete", d)

	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
	time.Sleep(500 * time.Millisecond)
	log.Println("-----")
}
