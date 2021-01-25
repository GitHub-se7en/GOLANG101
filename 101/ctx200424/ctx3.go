package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	// Set a duration.
	duration := 150 * time.Millisecond
	// Create a context that is both manually cancellable and will signal
	// cancel at the specified duration.
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// Create a channel to receive a signal that work is done.
	ch := make(chan string)

	// Ask the goroutine to do some work for us.
	go func(ch chan string) {
		defer waitGroup.Done()
		// Simulate work.
		time.Sleep(250 * time.Millisecond)

		// Report the work is done.
		ch <- "123"

	}(ch)

	// Wait for the work to finish. If it takes too long, move on.
	select {
	case d := <-ch:
		fmt.Println("work complete", d)

	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
	waitGroup.Wait()
	log.Println("-----")
}
