package main

import "con-system/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQRouting("exTest", "one")
	rabbitmq.RecieveRouting()
}