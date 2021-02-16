package main

import "con-system/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQRouting("exTest", "two")
	rabbitmq.RecieveRouting()
}