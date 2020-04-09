package main

import "fmt"

type notifier interface {
	notify()
}
type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("发送用户邮件给 %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

//func (ad *admin) notify() {
//	fmt.Printf("发送管理员邮件给 %s<%s>\n",ad.name,ad.email)
//}

func main() {
	ad := admin{
		user: user{
			name:  "weiguozhao",
			email: "weiguozhao@mywiz.com",
		},
		level: "super",
	}
	//给admin用户发送一个通知
	//用于实现接口的内部类型的方法，被提升到外部类型
	sendNotification(&ad)
	ad.user.notify()
	ad.notify()
}
func sendNotification(notifier2 notifier) {
	notifier2.notify()
}
