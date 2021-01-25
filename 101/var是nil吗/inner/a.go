package inner

import "log"

func A() {
	type J struct {
	}
}
func C(s struct{}) {
	type J struct {
	}
	log.Println(s)
}
