package main

import (
	"log"
	"strings"
)

const test = "testUhppp"

func main() {
	upper := strings.ToUpper(test)
	test = strings.ToUpper(test)
	replace := strings.Replace(test, "test", "123", 1)
	log.Println(upper)
	log.Println(replace)
}
