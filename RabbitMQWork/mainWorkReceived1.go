package main

import "con-system/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewSimpleRabbitMQ("simpleQueue")
	rabbitmq.ConsumeSimple()
}
