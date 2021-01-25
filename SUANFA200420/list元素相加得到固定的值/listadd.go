package main

import (
	"GOLANG101/SUANFA200420/common"
	"github.com/prometheus/common/log"
)

/**
- Input: list=[]int{2, 3, 4}, length=6
  Output: true, because there exists 2 and 4 that add up to 6
*/
/**
这个今天早上是突然间想到的，那就是每一步记录了减去的数，假如说，我按照最正常的逻辑进行遍历，首先是得到第一个，然后用总数减去这个数，同时记录下这个数，
用减法是我没有想到的，
按照我的理解的话，这个必须遍历两次才行，但是人家就是用一次遍历就完成了所有的结果，怎么做到的？
这种思想，我很难理解，这个是全部遍历的，如果是两遍遍历的话，是会停止循环的
肯定是要走回头路的，这就跟将一根棍子进行对折一样，前半部分，只记录，后半部分才取出来
*/
func main() {
	test := []struct {
		arr      []int
		totalNum int
		expected bool
	}{
		{[]int{}, 1, false},
		{[]int{0}, 1, false},
		{[]int{1}, 1, true},
		{[]int{0, 1}, 1, true},
		{[]int{1, 1}, 2, true},
		{[]int{2, 3, 4}, 6, true},
		{[]int{2, 3, 4}, 8, false},
	}

	for key, value := range test {
		log.Info(key)
		result := countNum(value.arr, value.totalNum)
		common.Equal(value.expected, result)
	}
}
func countNum(numArray []int, totalLength int) bool {
	rememberNum := make(map[int]int)

	for _, value := range numArray {
		need := totalLength - value
		if _, ok := rememberNum[need]; ok {
			return true
		}
		rememberNum[value] = 1
	}

	return false
}
