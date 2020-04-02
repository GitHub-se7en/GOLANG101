package main

func main() {
	//i := GetValue()
	i := GetValue2()

	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}

func GetValue() int          { return 1 }
func GetValue2() interface{} { return 1 }
