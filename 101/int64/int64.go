package main

import "fmt"

func main() {
	var r int8
	//int8代表的是8个字节，一个字节可以放一个数字
	//能推导出，放8个数字吗？
	r = -129
	fmt.Print(r)
}
