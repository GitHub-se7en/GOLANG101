package main

import "fmt"

func main() {
	//修改字符串的第一个字母，这个做法之所以可以行，完全是因为go里面字符串就是byte数组组成的
	s := "hello"
	c := []byte(s) // 将字符串 s 转换为 []byte 类型
	c[0] = 'c'
	s = string(c)       // 再转换回 string 类型
	fmt.Printf("%s", s) // cello
}
