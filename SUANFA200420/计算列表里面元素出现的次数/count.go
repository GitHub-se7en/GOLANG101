package main

import "log"

func main() {
	in := []int{2, 3, 4, 2, 6, 7, 8}
	counts := make([]int, 9)
	for _, item := range in {
		counts[item]++
		//i := counts[item]
		//log.Println(i)
		//counts[item] = i+1
	}
	log.Println(counts)
}
