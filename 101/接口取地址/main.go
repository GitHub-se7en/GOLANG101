package main

type asd interface {
}
type acb struct {
	Id int
}

func main() {
	//i2 := &c()
	//i := &d()
	//i3 := &acb{}
	acb{}.f()

}
func (*acb) f() {

}
func c() asd {
	return nil
}
func d() *acb {
	return nil
}
