package main

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// 1. Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Connection failed:", err)
	}
	defer conn.Close()

	// 2. Create channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Channel error:", err)
	}
	defer ch.Close()

	// 3. Declare queue
	q, err := ch.QueueDeclare(
		"email_queue", // queue name
		true,          // durable
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	// 4. Publish message
	message := "StudentCreated:101"

	err = ch.PublishWithContext(
		context.Background(),
		"",     // default exchange
		q.Name, // routing key = queue name
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		log.Fatal("Publish failed:", err)
	}

	log.Println("Message sent :", message)
}
