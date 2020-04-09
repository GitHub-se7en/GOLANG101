package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!")
	fmt.Fprint(os.Stdout, "world!")
	//w.Flush() // Don't forget to flush!
}
