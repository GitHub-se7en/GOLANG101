package RPC

import (
	"fmt"
	"testing"
)

func TestZhizhen(t *testing.T) {
	var i int = 5
	a(&i)
	fmt.Print(i)
}
func a(i *int) *int {
	*i = 10
	return i
}
