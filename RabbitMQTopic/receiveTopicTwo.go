package main

import "con-system/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQTopic("testTopic", "test.*.two")
	rabbitmq.ReceiveTopic()
}
