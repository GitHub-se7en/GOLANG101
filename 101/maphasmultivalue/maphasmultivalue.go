package main

import "log"

func main() {
	strings := make(map[string]string)
	strings["zhaoweiguo"] = "我是赵卫国的value"
	strings["zhaoweiguo"] = "我是赵卫国的value2"
	strings[&"zhaoweiguo"] = "我是赵卫国的value2"
	zhaoweiguobianliang := "zhaoweiguo"
	strings[&zhaoweiguobianliang] = "我是赵卫国的value2"
	s := strings["zhaoweiguo"]
	log.Println(s)
}
