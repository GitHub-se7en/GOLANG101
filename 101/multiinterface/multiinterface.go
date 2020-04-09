package main

import "fmt"

type abs interface {
	bs()
}
type A struct{}
type B struct{}

func (b B) bs() {
	fmt.Print("我是B实现abs的方法")
}
func (a A) bs() {
	fmt.Print("我是A实现abs的方法")
}

func main() {
	a := abs(B{})
	a.bs()
}
