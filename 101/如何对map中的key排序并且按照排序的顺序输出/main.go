package main

import (
	"fmt"
	"sort"
)

func main() {

	var m = map[string]int{
		"hello":   0,
		"morning": 1,
		"my":      2,
		"girl":    3,
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		//厉害的地方其实就是这个通过key获取value的方法,取出key来,然后排序,然后根据key获取
		fmt.Println("Key:", k, "Value:", m[k])
	}
}
