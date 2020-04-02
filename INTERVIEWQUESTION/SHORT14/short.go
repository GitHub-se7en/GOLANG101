package main

import "fmt"

var (
	size     = 1024
	max_size = size * 2
)

const (
	x = iota
	y = "zz"
	k = iota
)

func main() {
	fmt.Println(x, y, k)
}
func main1() {
	println(size, max_size)
}
