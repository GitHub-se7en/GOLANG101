package main

import "fmt"

func main() {
	list := new([]int)
	list1 := []int{}
	//list = append(list, 1)
	list1 = append(list1, 1)
	fmt.Println(list)
	fmt.Println(list1)
}
