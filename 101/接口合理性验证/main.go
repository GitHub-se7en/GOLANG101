package main

import (
	"log"
	"net/http"
)

type Handler struct {
	// ...
}
type asd struct {
}

func main() {
	a := (*asd)(nil)
	log.Println(a)
	//这个是匿名变量,空格之后是类型,  并不是之前的下划线加上变量表示不使用
	var _ http.Handler = (*Handler)(nil)
}
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ...
}
