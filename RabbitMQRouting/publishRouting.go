package main

import (
	"con-system/RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	rabbitmq1 := RabbitMQ.NewRabbitMQRouting("exTest", "one")
	rabbitmq2 := RabbitMQ.NewRabbitMQRouting("exTest", "two")
	for i := 0; i < 10; i++ {
		rabbitmq1.PublishRouting("Hello RabbitMQ Routing one" + strconv.Itoa(i))
		rabbitmq2.PublishRouting("Hello RabbitMQ Routing two" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
