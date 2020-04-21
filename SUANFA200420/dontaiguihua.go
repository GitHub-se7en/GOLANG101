package main

import (
	"log"
)

//递归，动态规划，记录已经算过的数据，两个变量
//国王和金矿问题
func main() {
	goldHas := []int64{350, 400, 500, 200, 300}
	goldNeedPeople := []int64{3, 5, 5, 3, 4}
	result := getMaxGold(10, 5, goldHas, goldNeedPeople)
	log.Println(result)
}

//如果是使用二维数组记录下来数据，不在重复计算
func getMaxGold(peopleNum int64, goldNum int64, goldHas []int64, goldNeedPeople []int64) int64 {
	if peopleNum < goldNeedPeople[0] && goldNum == 1 {
		return 0
	} else if goldNum == 1 {
		return goldHas[0]
	}
	maxGold := getMaxGold(peopleNum, goldNum-1, goldHas, goldNeedPeople)
	if peopleNum < goldNeedPeople[goldNum-1] {
		return maxGold
	}
	maxGold2 := getMaxGold(peopleNum-goldNeedPeople[goldNum-1], goldNum-1, goldHas, goldNeedPeople) + goldHas[goldNum-1]
	return Max(maxGold, maxGold2)
}

func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}
