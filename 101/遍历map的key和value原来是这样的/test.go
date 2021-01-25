package main

import "log"

func main() {
	bools := map[string]bool{"cat": true, "cta": true, "act": true, "atc": true, "tca": true, "tac": true}
	for key, value := range bools {
		log.Println(key, "-----", value)
	} //原来遍历map得到的key就是key
	/**
	2020/06/11 10:28:50 cat ----- true
	2020/06/11 10:28:50 cta ----- true
	2020/06/11 10:28:50 act ----- true
	2020/06/11 10:28:50 atc ----- true
	2020/06/11 10:28:50 tca ----- true
	2020/06/11 10:28:50 tac ----- true
	*/
}
