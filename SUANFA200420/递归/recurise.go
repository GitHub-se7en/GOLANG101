package main

import (
	"GOLANG101/SUANFA200420/common"
	"strings"
)

/**
啥意思啊，为什么是只留下最后一个，把之前的进行所有的排列，然后把最后一个数，根据不同的位置放到所有的组合里面呢？
给定一个字符串，输出所有的排列，不能重复
排列在计算机里面是怎么实现的呢？？
如果是按照阶乘的方式实现的话,
按照计算机的做法的话，这个的实现就是先固定一个，然后插空，一个产生两个空
最原始的排列组合还真是这样，难道是后来数学里面总结出来了规律了
这两个正好是倒着来的，一个固定了格子，不断变化数据，另外一个是固定了数据，不断变化格子，两种完全是不一样的解法
*/
/**
递归实现的排列组合，递归的结束点是只有一个数的时候，然后不断累加
*/
func main() {
	tests := []struct {
		in       string
		expected map[string]bool
	}{
		{"cat", map[string]bool{"cat": true, "cta": true, "act": true, "atc": true, "tca": true, "tac": true}},
	}

	for _, tt := range tests {
		result := getAll(tt.in)
		common.Equal(tt.expected, result)
	}
}
func getAll(in string) map[string]bool {
	results := make(map[string]bool)
	if len(in) <= 1 {
		results[in] = true
		return results
	}
	lastChar := string(in[len(in)-1])
	beforeLastChar := string(in[:len(in)-1])

	allSort := getAll(beforeLastChar)
	//把最后一个字符插入到所有的排列里面的所有的空格里面
	for mapKey, _ := range allSort {
		//我还没办法用map里面的数据，没办法获取key啊，，尴尬了吧，直接遍历就能得到map的key，傻了吧
		//这个地方没有按照书上写的，而是按照自己的理解，我用的是mapKey这个变量，然后在每一个空里面插入最后一个字符
		for key, _ := range mapKey {
			s := []string{mapKey[:key], lastChar, mapKey[key:]}
			join := strings.Join(s, "")
			results[join] = true
			if key == 0 {
				s := []string{mapKey, lastChar}
				p := strings.Join(s, "")
				results[p] = true
			}

		}
	}
	return results
}
