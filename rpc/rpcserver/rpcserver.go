package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct{}
type RpcServer int64

func (t *RpcServer) ProvideServerTime(args *Args, reply *int64) error {
	*reply = time.Now().Unix()
	return nil
}

func main() {
	// Initialize new RPC server
	rpcServer := new(RpcServer)
	// Register RPC
	rpc.Register(rpcServer)
	rpc.HandleHTTP()
	// Open and Listen on port 9000
	l, e := net.Listen("tcp", ":9000")
	if e != nil {
		log.Fatal("listener error:", e)
	}

	http.Serve(l, nil)
}
