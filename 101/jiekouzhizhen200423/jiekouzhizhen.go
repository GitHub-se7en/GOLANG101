package main

import "log"

type a interface {
	p()
}

type b struct {
}

func (b b) p() {

}
func main() {
	i := &a(b{})
	l := b{}
	log.Println(i, l)
}
