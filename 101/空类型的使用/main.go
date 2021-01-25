package main

import "log"

func main() {
	//空类型
	i := struct {
	}{}
	log.Println(i)
}
