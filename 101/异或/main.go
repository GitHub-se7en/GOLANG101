package main

import "log"

//异或,自己和自己异或会是0
func main() {
	i := 5
	j := 5
	i ^= j
	log.Println(i)
	k := 3
	p := 4
	result := k & p
	log.Println(result)
}
