package main

import "log"

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Println("err不为nil", err)
		} else {
			log.Println("err为nil")
		}
		log.Println("我是defer")
	}()
	panic("我是panic")
}
