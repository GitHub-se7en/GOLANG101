package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/net/context"
)

var (
	wg sync.WaitGroup
)

func work(ctx context.Context) error {
	defer wg.Done()
	//这个并不是同时启动1000个，这个是执行第一个的时候就卡在那等待了
	for i := 0; i < 1000; i++ {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Doing some work ", i)

		// we received the signal of cancelation in this channel
		case <-ctx.Done():
			fmt.Println("Cancel the context ", i)
			return ctx.Err()
		}
	}
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	fmt.Println("Hey, I'm going to do some work")

	wg.Add(1)
	go work(ctx)
	wg.Wait()

	fmt.Println("Finished. I'm going home")
}
