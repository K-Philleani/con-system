package main

import (
	"con-system/RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewSimpleRabbitMQ("simpleQueue")

	for  i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Hello. RabbitMQ!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
