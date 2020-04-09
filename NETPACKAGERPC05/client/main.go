package main

import (
	"GOLANG101/NETPACKAGERPC05"
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	} // Synchronous call
	args := &awesomeProjectGORPC.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
	// Asynchronous call
	quotient := new(awesomeProjectGORPC.Quotient)
	var s awesomeProjectGORPC.Quotient
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	// check errors, print, etc.
	if replyCall.Error != nil {
		log.Fatal("divide error:", err)
	}
	fmt.Printf("Divide: %d--%d--%d", args, quotient.Quo, quotient.Rem)
}
