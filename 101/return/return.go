package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Print(Dsa())
}
func Afd() (string, error) {
	return "ss", errors.New("这个是错误")
}
func Dsa() (e error) {
	_, e = Afd()
	return
}
