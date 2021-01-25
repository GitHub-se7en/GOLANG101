package main

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	m sync.Mutex
	n uint64
}

func (c *counter) value() uint64 {
	c.m.Lock()
	defer c.m.Unlock()
	return c.n
}

func (c *counter) increase(delta uint64) {
	c.m.Lock()
	c.n += delta
	c.m.Unlock()
}
func main() {
	var c counter
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 1; i <= 100; i++ {
		go func() {
			defer wg.Done()
			for k := 1; k <= 100; k++ {
				time.Sleep(time.Millisecond)
				c.increase(1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(c.value())
}
