package main

import "log"

func main() {
	runes := []rune("世界")
	log.Println(runes)
	log.Println(string(byte(97)))
	log.Println(byte(97))
	log.Println(string(byte(90)))
	log.Println(string([]byte{12, 12, 12}))
	log.Println([]byte("是"))
	//strings := make(map[byte]int)
	//s := strings["A"]
	//log.Println(s)

}
