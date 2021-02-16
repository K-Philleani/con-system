package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// 订阅模式创建RabbitMQ实例
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ{
	// 创建RabbitMQ实例
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	var err error
	// 获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.FailOnErr(err, "failed to connect rabbitmq")
	// 获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.FailOnErr(err, "failed to open a channel")
	return rabbitmq
}


// 订阅模式生产
func (r *RabbitMQ) PublishPub(message string) {
	// 1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
		)
	r.FailOnErr(err, "Failed to declare an exchange")

	// 2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		})
}

// 订阅模式消费端
func (r *RabbitMQ) ReceiveSub() {
	// 1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.FailOnErr(err, "Failed to declare an exchange")

	// 2. 尝试创建队列
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
		)
	r.FailOnErr(err, "Failed to declare a queue")

	// 绑定队列到exchange
	err = r.channel.QueueBind(
		q.Name,
		"",
		r.Exchange,
		false,
		nil,
		)

	// 消费消息
	message, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)
	forever := make(chan bool)
	go func() {
		for d := range message {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	fmt.Printf("退出请按 CTRL+C\n")
	<-forever
}