package main

import "log"
import child "GOLANG101/101/包级别变量大小写问题/child"

func s() {
	aCopy := a
	i := child.Child
	log.Println(a1{})
	log.Println(a)
	log.Println(aCopy)
	log.Println(i)
}
