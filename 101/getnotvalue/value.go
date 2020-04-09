package getnotvalue

import "fmt"

// duration 是一个基于 int 类型的类型
type duration int

// 使用更可读的方式格式化 duration 值
func (d *duration) pretty() {
	fmt.Println("Duration:", *d)

}

// main 是应用程序的入口
func Value() {
	duration(42).pretty()
	//v :=duration(42)
	//&v.pretty()
	//fmt.Println("Duration:",&duration(42))
	// ./listing46.go:17: 不能通过指针调用 duration(42)的方法
	// ./listing46.go:17: 不能获取 duration(42)的地址
}
