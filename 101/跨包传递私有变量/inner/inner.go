package inner

func A() struct {
	id string
} {
	i := struct {
		id string
	}{
		"123",
	}
	return i

}
func B(s struct{}) {

}
func C(func()) {

}
