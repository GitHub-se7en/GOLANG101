//package innerClass
package main

import "fmt"

type User struct {
	Admin
	name  string
	email string
}
type Inter interface {
	NotifyAdmin() Inter
}

type Admin struct {
	u     *User //嵌入类型
	level string
}

func (ad *Admin) NotifyAdmin() Inter {
	fmt.Printf("我是admin的方法")
	return ad
}

//func main() {
//	admin := Admin{
//		u: &User{
//			name:  "weiguozhao",
//			email: "weiguozhao@mywiz.com",
//		},
//		level: "super",
//	}
//	admin.u.NotifyAdmin()
//	u := &User{
//		name:  "adminzhaoweiguo",
//		email: "adminzhaoweiguo@qq.com",
//	}
//	u.NotifyAdmin()
//}
