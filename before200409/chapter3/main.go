package main

import (
	"GOLANG101/before200409/chapter3/words"
	"fmt"
	"io/ioutil"
	"os"
)

// main是应用程序的入口
func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)

	fmt.Printf("这里有 %d 字符在你的文件里面。\n", count)

}
