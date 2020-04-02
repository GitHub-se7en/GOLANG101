package main

func test(x int) (func(), func()) {
	return func() {
		println(x)
		x += 10
	}, func() { println(x) }
}

func main1() {
	a, b := test(100)
	a()
	b()
}
