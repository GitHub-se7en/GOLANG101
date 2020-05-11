package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var f *os.File
	f, _ = os.Create("./PAOPI200508/跑批数据.txt") //创建文件
	defer f.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(f)
	for i := 1; i <= 15000; i++ {
		n, _ := write.WriteString(strconv.Itoa(i) + "\r\n")
		fmt.Printf("写入 %d 个字节n", n)
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
	f.Close()

}
