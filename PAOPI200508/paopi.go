package main

import (
	"GOLANG101/PAOPI200508/db"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	//先查询出来所有的数据，然后放到slice里面，然后遍历每一个开启一个goroutine，修改数据库
	//怎么连接mysql？不能使用gorm，太费劲了，使用别的项目

	readFileWithBuffer()
}
func readFile() {

	var wg sync.WaitGroup

	file, _ := os.Open(`跑批数据.txt`)
	defer file.Close()
	//有一个思路是创建两个channel，一个是存放错误信息使用的，一个是存放正确的信息的
	correctData := make(chan string)
	wrongData := make(chan string)

	//创建缓冲读取器
	reader := bufio.NewReader(file)
	for {

		//读取一行
		lineBytes, _, err := reader.ReadLine()

		//读取完毕时，关闭所有数据管道，并退出读取
		if err == io.EOF {
			fmt.Println("已经读到文件末尾！")
			break
		}
		wg.Add(1)
		//拿出省份ID
		id := string(lineBytes)

		go func(id string) {
			defer wg.Done()
			_, e := db.UpdateTime(id)
			if e != nil {
				wrongData <- id
			} else {
				correctData <- id
			}
		}(id)
	}
	go func() {
		wg.Wait()
		close(correctData)
		close(wrongData)
	}()

	go func() {
		for value := range correctData {
			log.Println(value)
		}
	}()
	for value := range wrongData {
		log.Println(value)
	}
}

func readFileWithBuffer() {

	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	log.Println("开始处理数据", time.Now().Format("2006-01-02 15:04:05"))
	file, err := os.Open(`./PAOPI200508/跑批数据.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//有一个思路是创建两个channel，一个是存放错误信息使用的，一个是存放正确的信息的
	correctData := make(chan string, 10000)
	wrongData := make(chan string, 10000)

	//创建缓冲读取器
	reader := bufio.NewReader(file)
	var count int
	for {

		//读取一行
		lineBytes, _, err := reader.ReadLine()
		count++
		//读取完毕时，关闭所有数据管道，并退出读取
		if err == io.EOF {
			fmt.Println("已经读到文件末尾！", count)
			break
		}
		if err != nil {
			log.Println(err)
			break
		}
		wg.Add(1)
		//拿出ID
		id := string(lineBytes)
		go func(id string) {
			defer wg.Done()
			_, e := db.UpdateTime(id)
			if e != nil {
				wrongData <- id
			} else {
				correctData <- id
			}
		}(id)
	}
	wg2.Add(2)
	go func() {
		defer wg2.Done()
		for {
			id, ok := <-wrongData
			if !ok {
				log.Println("wrongData空了")
				return
			} else {
				log.Println("wrongData", id)
			}
		}

	}()
	go func() {
		defer wg2.Done()
		for {
			id, ok := <-correctData
			if !ok {
				log.Println("correctData空了")
				return
			} else {
				log.Println("correctData", id)
			}
		}
	}()
	wg.Wait()
	close(correctData)
	close(wrongData)
	wg2.Wait()
	log.Println("处理完成数据", time.Now().Format("2006-01-02 15:04:05"))

}
