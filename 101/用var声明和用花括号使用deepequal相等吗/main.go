package main

import (
	"log"
	"reflect"
)

func main() {
	//这种方式称之为切片字面量声明切片，
	float64s1 := []float64{}
	//这个是声明空切片的两种方式，通过reflect 判断出来居然是true的
	float64s3 := make([]float64, 0)
	var float64s2 []float64
	log.Println(float64s1)
	log.Println(float64s2)
	log.Println(reflect.DeepEqual(float64s1, float64s2))
	log.Println(reflect.DeepEqual(float64s1, float64s3))
}
