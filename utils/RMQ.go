package utils

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

// RabbitMQ连接函数
func RabbitMQConn() (conn *amqp.Connection, err error) {
	// RabbitMQ分配的用户名称
	var user string = "admin"
	// RabbitMQ用户的密码
	var pwd string = "admin123"
	// RabbitMQ Broker 的ip地址
	var host string = "106.14.223.102"
	// RabbitMQ Broker 监听的端口
	var port string = "5672"
	url := "amqp://" + user + ":" + pwd + "@" + host + ":" + port + "/"
	// 新建一个连接
	conn, err = amqp.Dial(url)
	// 返回连接和错误
	return
}

// 错误处理函数
func ErrorHanding(err error, msg string) {
	if err != nil {
		//log.Fatalf("%s: %s", msg, err)
		Logger.Info("阿斯顿发送到发", msg, err)
	}
}

func Push(massage interface{}) {

	// 连接RabbitMQ服务器
	conn, err := RabbitMQConn()
	ErrorHanding(err, "Failed to connect to RabbitMQ")
	// 关闭连接
	defer conn.Close()
	// 新建一个通道
	ch, err := conn.Channel()
	ErrorHanding(err, "Failed to open a channel")
	// 关闭通道
	defer ch.Close()

	// 声明或者创建一个队列用来保存消息
	q, err := ch.QueueDeclare(
		// 队列名称
		"simple:queue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	ErrorHanding(err, "Failed to declare a queue")

	//发送数据  json 字符串
	dataBytes, err := json.Marshal(massage)
	if err != nil {
		ErrorHanding(err, "struct to json failed")
	}
	Logger.Info("阿斯顿发送到发", err)

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        dataBytes,
		})
	log.Printf(" [x] Sent %s", dataBytes)
	ErrorHanding(err, "Failed to publish a message")
}
