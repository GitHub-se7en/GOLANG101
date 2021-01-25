package main

import (
	"log"
	"time"
)

func main() {
	go func() {
		defer func() {
			//err := recover()
			//if err != nil {
			//	log.Println("err不为nil", err)
			//} else {
			//	log.Println("err为nil")
			//}
			log.Println("我是defer")
		}()
		panic("我是panic")
	}()
	for {
		time.Sleep(time.Second)
		log.Println("--")
	}
}
