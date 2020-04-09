package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

//这个地方之所以使用指针，是因为我需要相同的request，而不是一个副本
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World,%s!", r.URL.Path[1:])
}
