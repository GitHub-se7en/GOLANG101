package main

import (
	"GOLANG101/GOROUTINEANDINTERFACE01/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}
func main() {
	search.Run("coronavirus")
}
