package main

import (
	"GOLANG101/SUANFA200420/common"
)

/**贪心算法   股票买卖
Example:
- Input: []int{10, 7, 5, 8, 11, 9}
  Output: 6, because one can buy at 5 and sell at 11
*/
/**
只可能是后面的比前面的大才能计算出结果，并不是找出最大最小值相减就完事的，只可能是后面减去前面，而不是前面减去后面，所以是按照时间线顺序排列，并且例子是买股票的例子
还是老样子，一次循环解决
用5减去前两个的最小值，就得到最大的利润
循环就代表着时间的流逝
*/
func main() {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{10}, 0},
		{[]int{10, 10}, 0},
		{[]int{10, 7, 5, 8, 11, 9}, 6},
		{[]int{10, 2, 5, 4, 7, 1}, 5},
		{[]int{10, 7, 2, 1}, -1},
	}

	for _, tt := range tests {
		result := timeCountMax(tt.in)
		common.Equal(tt.expected, result)
	}
}
func timeCountMax(stocks []int) int {
	if len(stocks) <= 2 {
		return 0
	}
	//第一个值就是最小值
	minPrice := stocks[0]
	//这里我本来想的是定义一个空值，但是不行，因为求max的时候会出现问题，所以人家定义的是重复值
	//var maxProfit int
	maxProfit := stocks[1] - stocks[0]
	//买和卖的顺序是一致的，肯定是先买后卖
	//而且这里必须是从1开始的
	for i := 1; i < len(stocks); i++ {
		currentPrice := stocks[i]
		potentialProfit := currentPrice - minPrice
		maxProfit = common.Max(potentialProfit, maxProfit)
		minPrice = common.Min(currentPrice, minPrice)
	} //循环过后，前面的就可以抛掉不用了啊，每次只计算两个
	//每循环一次就得出前面的list中的最小值和最大利润
	return maxProfit
}
