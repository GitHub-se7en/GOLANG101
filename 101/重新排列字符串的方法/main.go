package main

import (
	"log"
	"strings"
)

func main() {
	strings.Replace()
	s := "赵卫国"
	bytes := []byte(s)
	log.Println(bytes)
	log.Println(string(bytes))
	log.Println(s[1])
	log.Println(string(s[1]))
	log.Println(strings.Count(s, string(s[1])))
	log.Println(strings.Count(s, "赵"))
	log.Println(isRegroup("abc", "bac"))
}
func isRegroup(s1, s2 string) bool {
	sl1 := len([]rune(s1))
	sl2 := len([]rune(s2))

	if sl1 > 5000 || sl2 > 5000 || sl1 != sl2 {
		return false
	}

	for _, v := range s1 {
		log.Println("----------", strings.Count(s1, string(v)))
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}
