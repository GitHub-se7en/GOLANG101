package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
)

type Post struct {
	Id      string
	Content string
	Author  string
}

func store(data interface{}, filename string) {
	//创建缓冲区
	buffer := new(bytes.Buffer)
	//利用缓冲区创建encoder
	encoder := gob.NewEncoder(buffer)
	//将数据编码到缓冲区里面去，data到buffer
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}

}

func load(data interface{}, filename string) {
	//这个是数据的原始状态
	raw, e := ioutil.ReadFile(filename)
	if e != nil {
		panic(e)
	}
	//把这些数据装到缓冲区里面
	//buffer := bytes.NewBuffer(raw)
	//根据这些原始数据创建一个缓冲区，并借此为原始数据提供read，write方法
	buffer := new(bytes.Buffer)
	//创建buffer的目的是为了读写原始数据
	buffer.Write(raw)
	//为缓冲区创建相应的解码器
	decoder := gob.NewDecoder(buffer)
	//解密成具体的类型
	e = decoder.Decode(data)
	if e != nil {
		panic(e)
	}
}
func main() {
	post := Post{
		Id:      "1",
		Content: "你是谁",
		Author:  "赵卫国",
	}
	store(&post, "post1")
	var postRead Post
	load(&postRead, "post1")
	log.Println(postRead)
}
