package main

import (
	"bubble/utils"
	"log"
)

func main() {
	conn, err := utils.RabbitMQConn()
	utils.ErrorHanding(err, "failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	utils.ErrorHanding(err, "failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"simple:queue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	utils.ErrorHanding(err, "Failed to declare a queue")
	// 定义一个消费者
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.ErrorHanding(err, "Failed to register a consume")
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	select {}
}
