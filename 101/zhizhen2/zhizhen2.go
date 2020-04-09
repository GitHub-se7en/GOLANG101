package main

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

type serverCodec struct {
	sending sync.Mutex
	req     Request
	rwc     io.ReadWriteCloser
	dec     *gob.Decoder
	enc     *gob.Encoder
	encBuf  *bufio.Writer
	addr    net.Addr
	closed  bool
}

// Request is a header written before every RPC call. It is used internally
// but documented here as an aid to debugging, such as when analyzing
// network traffic.
type Request struct {
	Color         string        // color
	RemoteIP      string        // remoteIP
	Timeout       time.Duration // timeout
	ServiceMethod string        // format: "Service.Method"
	Seq           uint64        // sequence number chosen by client
}

func readRequest(codec *serverCodec) {
	var req = &codec.req
	log.Println(req)
	log.Println(*req)
	*req = Request{}
	log.Println(req)
	log.Println(*req)

}

func main() {
	srv := &serverCodec{}
	readRequest(srv)
}
