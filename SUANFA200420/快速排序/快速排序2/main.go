package main

/**
快速排序的思想
其实不光是快速排序,所有的算法基本上都是处理一个list而已
方法不局限与,取出来,
原地进行数组元素之间的替换



*/
func main() {
	quickSort([]int{5, 1, 1, 2, 2, 0, 0}, 0, 6)
}

func quickSort(input []int, start, end int) {
	if start < end {
		middle := getMiddle(input, start, end)
		quickSort(input, start, middle-1)
		quickSort(input, middle+1, end)
	}
}

func getMiddle(input []int, start int, end int) (middle int) {
	left := start
	right := end - 1
	for left <= right {
		for left <= end && input[left] < input[end] {
			left++
		}

		for right >= start && input[right] >= input[end] {
			right--
		}

		if left < right {
			temp := input[right]
			input[right] = input[left]
			input[left] = temp
		} else {
			temp := input[end]
			input[end] = input[left]
			input[left] = temp
		}
	}
	return left
}
