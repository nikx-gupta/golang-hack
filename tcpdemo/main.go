package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

func main() {
	st := time.Now()
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "google.com:80")
	fmt.Printf("Time Taken To Resolve: %s\n", time.Now().Sub(st))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", tcpAddr)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		panic(err)
	}

	result, err := ioutil.ReadAll(conn)
	fmt.Println(string(result))
}
