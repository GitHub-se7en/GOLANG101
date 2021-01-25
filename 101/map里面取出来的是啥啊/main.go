package main

import "log"

func main() {
	ints := make(map[rune]int)
	ints[32] = 32
	ints[33] = 33
	ints[34] = 34
	ints[35] = 35
	if v, ok := ints[33]; ok {
		log.Println(v, ok)
	}
	for i, ch := range ints {
		log.Println(i, ch)
	}
	s := "asdf"
	for key, value := range []byte(s) {
		log.Println(key, value)

	}
}
