package main

func main() {
	var array [5]int

	array1 := [5]int{10, 20, 30, 40, 50}

	array2 := [...]int{10, 20, 30, 40, 50, 60}

	array3 := [5]int{1: 20, 4: 40}

	a := &3
	a = new(int)
	a = &3
	*a = 3
}
