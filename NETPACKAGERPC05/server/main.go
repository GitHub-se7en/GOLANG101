package main

import (
	"GOLANG101/NETPACKAGERPC05"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	arith := new(awesomeProjectGORPC.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)

}
