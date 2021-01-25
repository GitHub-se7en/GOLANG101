package main

import "GOLANG101/SUANFA200420/common"

/**
Example:
- Input: 4
  Output: 3, because the 4th Fibonacci of the sequence [0, 1, 1, 2, 3] is 3

Approach:
- Instead of using a recursive approach, use a bottom-up approach and
  iteratively compute subsequent numbers until we get to n.
bottom uo方法原来是从最底层一步一步计算出来的，那递归呢
*/
/**
原来斐波那契数列数列是前两个数字相加，不断继续下去的数列
*/
func main() {
	tests := []struct {
		in       int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
	}

	for _, tt := range tests {
		result := fib(tt.in)
		common.Equal(tt.expected, result)
	}
}
func fib(n int) int {
	//目的是求第几个位置的斐波那契数列的数字是多少？
	//解决这个问题的办法就是从0开始计算到n-1
	//第几个位置和循环几次是有关系的啊
	current := 0
	f0 := 0
	f1 := 1

	if n == 0 || n == 1 {
		return n
	}
	//1到n-1我还是能计算出来这个一共是n-1个数的，但是0到n-2，我就算不出来了
	//要计算第n个位置的斐波那契数列，前面就需要计算n-1次相加
	for i := 1; i <= n-1; i++ {
		current = f0 + f1
		f0 = f1
		f1 = current
	}
	//f0和f1是不断进行变换的,交替前进的
	return current
}
