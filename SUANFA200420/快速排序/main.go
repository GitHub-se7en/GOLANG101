package main

import "log"

/**
1. 写了才知道原来人家取最后一个值作为待处理的值是有原因的,把最后一个值刨出去,剩下的队列,从左至右,从右至左,开始查找最后这个值应该放的位置,
这样的话,还解决了我之前的错觉,以为把中间这个值复制一份,它这个思路清晰 ,比我的清晰多了
2. 一开始不明白的返回left,原来是对的,因为按照这么推的话,那就是left是中间位置啊
*/
func main() {
	//什么是快速排序?
	//自己写出来才是硬道理,要不然看多少次都白搭,只不过是幻觉
	ints := []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	quick2(ints, 0, len(ints)-1)
	log.Println(ints)
}
func quick2(inputTotal []int, start, end int) {
	if start < end {
		middle := getMiddle(inputTotal, start, end)
		quick2(inputTotal, start, middle-1)
		quick2(inputTotal, middle+1, end)
	}
	//getMiddle(inputTotal, start, end)
}
func getMiddle(input []int, inputStart, inputEnd int) (middle int) {
	left := inputStart
	log.Println("左边的值为", input[left])
	right := inputEnd - 1
	log.Println("末尾减一的值为", input[right])
	middle = inputEnd
	log.Println("待处理的值为", input[middle])
	for left < right {
		for left < inputEnd && input[left] < input[middle] {
			log.Println("left", left, "middle", middle, "input[left]", input[left], "input[middle]", input[middle])
			left++
		}
		for right > inputStart && input[right] > input[middle] {
			log.Println("right", right, "middle", middle, "input[end]", input[right], "input[middle]", input[middle])
			right--
		}
		if left < right {
			tmp := input[right]
			input[right] = input[left]
			input[left] = tmp
		} else {
			temp := input[left]
			input[left] = input[middle]
			input[middle] = temp
		}
		log.Println(input)
	}

	return left
}
