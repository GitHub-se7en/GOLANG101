package main

import (
	"GOLANG101/JIEKOU/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}
func main() {
	search.Run("coronavirus")
}
