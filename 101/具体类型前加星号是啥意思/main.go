package main

import "log"

type abc struct {
	name string
}

func main() {
	a := abc{name: "zhaoweiguo"}
	v := &abc{name: "zhaoweiguo"}
	v.name
	var fa interface{}
	e, ok := fa.(*abc)
	log.Println(e, ok)

	log.Println(*abc)

}
