package main

import "fmt"

func main() {
	var array [5]int
	array1 := [5]int{10, 20, 30, 40, 50}
	array2 := [...]int{10, 20}
	array3 := [5]int{1: 10, 2: 20}
	fmt.Println(array[1])
	fmt.Println(array1[1])
	fmt.Println(array2[1])
	fmt.Println(array3[1])
	fmt.Println("------数组声明的几种方式--------")

	array4 := [5]*int{0: new(int), 1: new(int)}
	*array4[0] = 10
	*array4[1] = 20
	fmt.Println(*array4[0])
	fmt.Println("-----数组指针-------")

	var arrayCopy [2]string
	arrayCopy2 := [2]string{"red", "blue"}
	arrayCopy = arrayCopy2
	fmt.Println(arrayCopy)
	fmt.Println("-----数组复制，得到两个一毛一样的-------")

	var arrayCopyStar [2]*string
	arrayCopyStar2 := [2]*string{new(string), new(string)}
	*arrayCopyStar2[0] = "Red"
	*arrayCopyStar2[1] = "Black"
	arrayCopyStar = arrayCopyStar2
	fmt.Println(*arrayCopyStar[0])
	fmt.Println(*arrayCopyStar[1])
	fmt.Println("-----数组指针复制，只能得到一个-------")

}
