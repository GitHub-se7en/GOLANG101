package main

import "log"

/**
冒泡排序第一反应
两个for循环,像泡泡一样不断向上浮上去
*/
/**
如何选出一个最大值?
拿着这个和别人比-----选择排序
两个两个比,比较并交换?
*/
/**
三个人如何选出最高的那个人?
方案是两个人选出一个最高的就好
*/
func main() {
	ints := []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	bubbleSort3(ints)
	log.Println(ints)
}

func bubbleSort3(input []int) {
	for j := 0; j < len(input); j++ {
		//这层循环是计算的次数
		for i := 0; i < len(input)-1-j; i++ {
			if input[i] > input[i+1] {
				temp := input[i+1]
				input[i+1] = input[i]
				input[i] = temp
			}
		}
	}
}

//func bubbleSort2(input []int) {
//	for i := 0; i < len(input); i++ {
//		for j := i+1; j < len(input); j++ {
//
//		}
//	}
//}
//func bubbleSort(input []int) {
//	for key, value := range input {
//		for key, value := range input {
//
//		}
//	}
//}
