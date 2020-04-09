package main

import "log"

type A struct {
	name   string
	gender string
}
type B struct {
	A
	id string
}

func main() {
	a := A{
		name:   "赵卫国",
		gender: "男",
	}
	b := B{
		A:  a,
		id: "a",
	}
	log.Print(a)
	log.Println()
	log.Print(b)
}
