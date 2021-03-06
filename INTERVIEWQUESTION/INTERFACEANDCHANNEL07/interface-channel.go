package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

//下面的迭代会有什么问题？

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	//ch := make(chan interface{}) // 解除注释看看！
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()

		for elem, value := range set.s {
			ch <- elem
			println("Iter:", elem, value)
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func main() {

	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}
	th.Iter()
	//fmt.Sprintf("%s%v", "ch", v)

	const name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)
	io.WriteString(os.Stdout, s)
}
