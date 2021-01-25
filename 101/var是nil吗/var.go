package main

import (
	"GOLANG101/101/var是nil吗/inner"
	"log"
)

func main() {
	var i int
	log.Println(i)
	type A struct {
	}
	type a2 struct {
	}
	//inner.J{}
	inner.C(A{})
	inner.C(a2{})
}
