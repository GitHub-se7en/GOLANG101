package main

import (
	"gorm.io/datatypes"
	"log"
)
import (
	jsoniter "github.com/json-iterator/go"
)

func main() {
	bytes, e := jsoniter.Marshal(`["red","blue"]`)
	log.Println(string(bytes),e)
	var test datatypes

}
