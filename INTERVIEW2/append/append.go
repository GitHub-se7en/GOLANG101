package main

import (
	"fmt"
)

func main() {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	fmt.Println(str2)
	//切片的含义就在这里了，所谓切片其实就是切的底层数组，所以str2改变了之后，str1也随着变了
	str2[1] = "new"
	fmt.Println(str1)
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1)
}
