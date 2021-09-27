package common

import (
	"github.com/streadway/amqp"
	"log"
)

func connection() string {
	conn := "amqps://tsxshdze:yJ01G2bcQsUSvIgldv9PAGC27A64w6S2@chimpanzee.rmq.cloudamqp.com/tsxshdze"

	return conn
}

func CreateChannel() (*amqp.Connection, *amqp.Channel) {
	conn := CreateConnection()

	ch, err := conn.Channel()

	FailOnError(err, "Error in Channel")

	return conn, ch
}

func CreateConnection() *amqp.Connection {
	conn, err := amqp.Dial(connection())
	FailOnError(err, "Error Connecting Client")

	return conn
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
