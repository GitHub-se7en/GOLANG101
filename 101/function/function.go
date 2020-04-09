package function

import "fmt"

func AA(func(s string)) string {
	fmt.Print("我是AA方法")
}
func BB() {
	AA(func(s string) {
		fmt.Print(s)
	})
	AA(cxc)
}
func cxc(s string) {
	fmt.Print(s)
}
