package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// URL格式：amqp://账号:密码@rabbitmq服务器地址:端口号/vhost
const MQURL = "amqp://admin:admin@124.70.71.78:5672/test"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机名
	Exchange string
	// Key
	Key string
	// 连接信息
	Mqurl string
}

// 创建RabbitMQ实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ{
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}
	// 创建连接
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.FailOnErr(err, "创建连接失败")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.FailOnErr(err, "获取channel失败")
	return rabbitmq
}

// 断开channel和connection的连接
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理
func (r *RabbitMQ) FailOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}