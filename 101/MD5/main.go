package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	str := "a"
	data := []byte(str)
	has := md5.Sum(data)
	md5str16 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	fmt.Println(md5str16)
}
