package main

import (
	"devignite/rabbitmq/client"
	"devignite/rabbitmq/server"
	"fmt"
)

func main() {

	fmt.Println("Executing")
	SimpleQueue()
}

func SimpleQueue() {
	//go func() {
		server.PublishSimpleQueue()
	//}()

	//go func() {
		client.SimpleQueue()
	//}()
}
