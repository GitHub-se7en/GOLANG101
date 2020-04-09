package value

import "fmt"

func Aaaa() {
	var s string = "ssss"
	var d = 34
	fmt.Print(s)
	fmt.Print(d)
	s1 := "ssss"
	d1 := 34
	fmt.Print(s1)
	fmt.Print(d1)
	//下面这行编译会报错的，不能把struct里面的格式声明变量
	//s4 string = "sdf"
	//s2 = 34
	//fmt.Print(s2)
}
