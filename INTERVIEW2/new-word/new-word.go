package main

import "fmt"

type passport struct {
	Photo       []byte
	Name        string
	Surname     string
	DateOfBirth string
}

func main() {
	p1 := passport{}
	var p2 passport
	p3 := passport{
		Photo:       make([]byte, 0, 0),
		Name:        "Scott",
		Surname:     "Adam",
		DateOfBirth: "Some time",
	}

	fmt.Printf("%s\n%s\n%s\n", p1, p2, p3)
	fmt.Println("--------------")

	pointerp1 := &p3
	fmt.Printf("%s", pointerp1)
	fmt.Println("--------------")
	pointerp2 := new(passport)
	pointerp2.Name = "Anotherscott"
	fmt.Printf("%s", pointerp2)
}
