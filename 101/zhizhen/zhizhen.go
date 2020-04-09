package zhizhen

import "fmt"

type Duration int

func GetAddress() {
	strings := [2]*string{new(string), new(string)}
	*strings[0] = "Red"
	*strings[1] = "Blue"
	fmt.Println(*strings[0])
	fmt.Println(&strings)
	fmt.Println("----------------------")
	var s string
	//*s = "sss"

	fmt.Println(s)
	//这个指的是指针所指向的值
	fmt.Println(*s)
	//取得是这个值的地址
	fmt.Println(&s)
	fmt.Println("----------------------")
	dur := Duration(64)
	fmt.Println(&dur)
}
