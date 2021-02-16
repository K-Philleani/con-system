package main

import "con-system/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("newProduct")
	rabbitmq.ReceiveSub()
}