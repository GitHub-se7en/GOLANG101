package main

import (
	"GOLANG101/101/defer和recover居然是协程级别的啊啊啊啊啊/child"
	"github.com/prometheus/common/log"
	"time"
)

func main() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				//r = errors.WithStack(r.(error))
				log.Error("recover panic  error(%+v)", r)
			}
		}()
		child.Count()
	}()
	time.Sleep(time.Second * 5)

}
