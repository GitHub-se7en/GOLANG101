package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("++++")
			f := err.(func() string)
			s := reflect.TypeOf(err).Kind().String()
			fmt.Println(s)
			fmt.Println(err, f())
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic(func() string {
			return "defer panic  1\n"
		})
	}()
}
