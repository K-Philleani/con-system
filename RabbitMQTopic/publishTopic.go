package main

import (
	"con-system/RabbitMQ"
	"log"
	"strconv"
	"time"
)

func main() {
	testOne := RabbitMQ.NewRabbitMQTopic("testTopic", "test.topic.one")
	testTwo := RabbitMQ.NewRabbitMQTopic("testTopic", "test.topic.two")
	for i := 0; i<10; i++{
		testOne.PublishTopic("Topic模式One第" + strconv.Itoa(i) +"条数据")
		testTwo.PublishTopic("Topic模式Two第" + strconv.Itoa(i) +"条数据")
		log.Println("Topic模式生产第" + strconv.Itoa(i) +"条数据")
		time.Sleep(1 * time.Second)
	}
}
