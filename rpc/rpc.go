package main

import (
	"log"
	"net/rpc"
)

type Args struct {
}

func main() {
	var reply int64
	args := Args{}
	client, err := rpc.DialHTTP("tcp", "localhost:9000")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	err = client.Call("RpcServer.ProvideServerTime", args, &reply)
	if err != nil {
		log.Fatal("Receiving Error:", err)
	}

	log.Printf("%d", reply)
}
