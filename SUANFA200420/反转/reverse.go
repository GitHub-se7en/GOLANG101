package main

import "GOLANG101/SUANFA200420/common"

/**
- Input: []string{"w", "o", "r", "l", "d", "", "h", "e", "l", "l", "o", "", "s", "a", "y"}
  Output: []string{"s", "a", "y", "", "h", "e", "l", "l", "o", "", "w", "o", "r", "l", "d"}
say hello world
这种情况下的解题思路就是先进行一次全部反转，得到正确的单词顺序，但是每个单词里面的字母的顺序是不一致的，所以针对每一个在进行一次反转
*/
/**
如果是我在做这个题目的时候，我就会切分出来，然后进行反转，因为这样不需要改动已经好的东西，但是按照我这样操作的话，更麻烦了，
人家这个就巧妙在只用了交换这一个功能，重复做，简单事情重复做，算法的魅力
答案就是reverse两次，从这个答案可以推出什么
这个题目的迷惑性就是让人不想动已经排好序的部分，这样就会导致一直是错误的
中间的形成部分就是每隔一个空字符串就可以局部反转
这种题目没法弄，正着根本推不出来，倒着也推不出来，
问题，答案，方法，反转
如果按照最底层的划分，比较大小，那这一个的话就是list中的两个数调换位置
*/
func main() {
	test := []struct {
		in       []string
		expected []string
	}{
		{
			[]string{"w", "o", "r", "l", "d", "", "h", "e", "l", "l", "o", "", "s", "a", "y"},
			[]string{"s", "a", "y", "", "h", "e", "l", "l", "o", "", "w", "o", "r", "l", "d"},
		},
	}
	for _, value := range test {
		list := reverseWord(value.in)
		common.Equal(value.expected, list)
	}
}
func reverseWord(list []string) []string {
	//先整体所有的进行反转
	reverseChar(list, 0, len(list)-1)
	start := 0
	for i, _ := range list {
		if list[i] == "" {
			reverseChar(list, start, i-1)
			start = i + 1
		}
		if i == len(list)-1 {
			reverseChar(list, start, i)
		}
	}
	return list
}

//工具包里面是交换一个list里面的两个数据，但是全部反转怎么实现
func reverseChar(list []string, start int, end int) {
	for start < end {
		common.Swap(list, start, end)

		start++
		end--
	}
}
