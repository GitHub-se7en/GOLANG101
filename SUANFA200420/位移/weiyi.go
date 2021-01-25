package main

import "log"

func main() {
	findUniqueID([]int{1, 1, 3})
}

//异或？？？
//怎样找出唯一一个奇数次出现的元素
func findUniqueID(list []int) int {
	ones := 0

	for _, value := range list {
		ones ^= value
		log.Println(ones)
	}
	return ones
}
