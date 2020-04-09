package notify

import "fmt"

// notifier is an interface that defined notification
// type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

//实现了这个方法就代表着你这个类型是我这个类型的一种形态，也就是多态，
//go里面方法调用有时候是失败的，假设是使用地址绑定的方法，但是我是使用值去调用这个方法的，这样就会出现问题，go自己会获取这个值的地址，
// 但是获取不到，而且编译的时候是不会报错的，运行的时候就会报错
// notify implements a method with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the application.
func Notify() {
	// Create a value of type User and send a notification.
	u := user{"Bill", "bill@email.com"}

	u.notify()
	sendNotification(u)

	// ./listing36.go:32: cannot use u (type user) as type
	//                     notifier in argument to sendNotification:
	//   user does not implement notifier
	//                          (notify method has pointer receiver)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
