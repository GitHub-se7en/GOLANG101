package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal()
	}
	log.Println(name)
	cluster := ""
	zone := name
	split := strings.Split(name, "-")
	if len(split) >= 2 {
		zone = split[0]
		cluster = split[1]
	}
	log.Println(zone, cluster)

}
