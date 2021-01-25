package main

import (
	"log"
)

func main() {
	test := "a"
	for key, _ := range test {
		log.Println(key)
		log.Println(test[0:])
	}
	/**
	2020/06/10 17:51:05 0
	2020/06/10 17:51:05 1
	2020/06/10 17:51:05 2
	2020/06/10 17:51:05 3
	2020/06/10 17:51:05 4
	2020/06/10 17:51:05 5
	*/
}
