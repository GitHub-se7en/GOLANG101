package words

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/goinaction/code/chapter3/words"
)

func Words() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("这里有 %d 个单词在你的文件中。\n", count)
}
