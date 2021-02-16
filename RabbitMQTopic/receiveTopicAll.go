package main

import "con-system/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQTopic("testTopic", "#")
	rabbitmq.ReceiveTopic()
}
