package HANDLER03

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestEngine(t *testing.T) {
	engine := New()
	log.Println(engine.address)
	engine.addRoute("GET", "/add", func(ctx *Context) {
		log.Println("这个是handler1")
	}, func(ctx *Context) {
		log.Println("这个是handler2")
	})
	engine.Start()

	resp, err := http.Get("http://localhost:8080/add")
	if err != nil {
		t.Errorf("HTTPServ: get error(%v)", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("http.Get get error code:%d", resp.StatusCode)
	}
	fmt.Print(resp.Body)
	resp.Body.Close()

}
