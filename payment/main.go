package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
)

func main() {
	app := fiber.New()

	initRabbitMQ()

	consumeMessages()

	app.Listen(":3004")
}

func initRabbitMQ() {
	// Connect to RabbitMQ and initialize rabbitMQChannel
}

func consumeMessages() {
	var ch *amqp.Channel
	msgs, err := ch.Consume(
		"order_placed_queue", // queue
		"",                   // consumer
		true,                 // auto-ack
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // args
	)
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	go func() {
		for msg := range msgs {
			processPayment(msg.Body) // Implement this function to process payment
		}
	}()
}

func processPayment(body []byte) {
	// Parse the message body and process payment
	log.Printf("Received message: %s", body)
	// Implement your payment processing logic here
}
