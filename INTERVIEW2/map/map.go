package main

import (
	"fmt"
)

type Student struct {
	Age int
}

func main() {
	kv := map[string]Student{"menglu": {Age: 21}}
	//kv["menglu"].Age = 22
	student, exist := kv["menglu"]
	fmt.Println(student, exist)
	s := []Student{{Age: 21}}
	s[0].Age = 22
	fmt.Println(kv, s)
}
