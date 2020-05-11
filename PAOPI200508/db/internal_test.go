package db

import (
	"log"
	"testing"
)

func TestUpdateTime(t *testing.T) {
	err, _ := UpdateTime("1234")
	log.Println(err)
}
