package RPC

import (
	"fmt"
	"log"
	"testing"
)

type user struct {
	name string
}

//go语言里面的复制居然是真的复制，java里面的却不是
func TestFuzhi(t *testing.T) {
	user := user{name: "赵卫国"}
	fuzhi(user)
	log.Println(user.name)
}
func fuzhi(user2 user) {
	user2.name = "赵卫国复制"
}
func TestZhizhen(t *testing.T) {
	var i int = 5
	a(&i)
	fmt.Print(i)
}
func a(i *int) *int {
	*i = 10
	return i
}
