package main

//一个数组超过一半以上都是同一个数,求出这个数
func main() {

}

//一次循环一定是有一次的用处的,要么就是成为候选人,要么就干掉对方一滴血,要么给对方加上一滴血
func findMostNum(nums []int) int {
	candidate := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
			count = 1
		} else {
			if candidate == nums[i] {
				count++
			} else {
				count--
			}
		}
	}
	return candidate
}
