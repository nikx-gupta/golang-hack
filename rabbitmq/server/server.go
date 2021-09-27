package server

import (
	"devignite/rabbitmq/common"
	"fmt"
	"github.com/streadway/amqp"
)

func PublishSimpleQueue() {
	conn, ch := common.CreateChannel()

	defer ch.Close()
	defer conn.Close()

	q, err := ch.QueueDeclare(
		"simple", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)

	common.FailOnError(err, "Queue Declaration Failed ....")

	body := "Hello World"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	common.FailOnError(err, "Message Publishing Failed .....")

	fmt.Println("Publishing Message Complete")
}
