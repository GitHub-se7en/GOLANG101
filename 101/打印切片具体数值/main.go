package main

import (
	"fmt"
)

type Project struct {
	name string
}
type Proj struct {
	name string
}

func (p *Project) String() string {
	return "pname=" + p.name
}

func main() {
	projs := []*Proj{&Proj{"p1"}, &Proj{"p2"}}
	fmt.Println("projs: ", projs)
	fmt.Printf("projs: % +v\n", projs)
	fmt.Printf("projs: %v\n", projs)

	projects := []*Project{&Project{"p1"}, &Project{"p2"}}
	fmt.Println("projects: ", projects)
	fmt.Printf("projects: % +v\n", projects)
	fmt.Printf("projects: %v\n", projects)
}
