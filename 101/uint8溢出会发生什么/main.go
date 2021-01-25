package main

import "log"

/**
2020/10/16 14:27:56 255
2020/10/16 14:27:56 0
2020/10/16 14:27:56 1
*/
func main() {
	var i uint8
	for i = 0; i <= 255; i++ {
		log.Println(i)
	}
}
