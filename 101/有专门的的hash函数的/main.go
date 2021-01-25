package main

import (
	"fmt"
	"hash/fnv"
)

func main() {

	fmt.Println(hash("HelloWorld"))
	fmt.Println(hash("HelloWorld."))
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	var m map[string]string
	m["sdfa"] = "sdfads"
	fmt.Println(m)
	return h.Sum32()
}
