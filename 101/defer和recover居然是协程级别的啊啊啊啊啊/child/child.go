package child

import "log"

func Count() {
	log.Println("我是count函数,我是count函数")
	panic("count函数出现panic")
}
