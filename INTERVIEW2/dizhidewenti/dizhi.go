package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
}

func main() {
	fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"})
	fmt.Println(reflect.DeepEqual(&Student{Name: "menglu"}, &Student{Name: "menglu"}))
	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})
}
