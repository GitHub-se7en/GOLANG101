package main

import "log"

func main() {
	log.Println(lengthOfNonRepeatingSubStr("hello 世界!"))
}

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccuredMap := make(map[rune]int)
	//连java中的增强for循环都记不住,怎么记得住golang中的,这不是双标吗
	start := 0
	maxLength := 0
	for key, value := range []rune(s) {
		//记录的是上一次发生的位置,本次发生的位置是lastValue
		if lastValue, ok := lastOccuredMap[value]; ok && lastValue > start {
			start = lastOccuredMap[value] + 1
		}

		//计算长度
		if key-start+1 > maxLength {
			maxLength = key - start + 1
		}

		//存起来,value就是map的键,key就是map的值
		lastOccuredMap[value] = key
	}
	return maxLength
}
