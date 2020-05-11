package readLine_test

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestReadFileWithBuffer(t *testing.T) {
	file, e := os.Open(`../跑批数据.txt`)
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {

		//读取一行
		lineBytes, _, err := reader.ReadLine()

		//读取完毕时，关闭所有数据管道，并退出读取
		if err == io.EOF {
			fmt.Println("已经读到文件末尾！")
			break
		}
		if err != nil {
			log.Println(err)
			break
		}
		//拿出ID
		lineStr := string(lineBytes)
		log.Println(lineStr)
		//fieldsSlice := strings.Split(lineStr, ",")
		//这个看不懂，反正是取到某一个特殊位置的数据作为id
		//id := fieldsSlice[1][0:2]
		//log.Println(id)
	}

}
