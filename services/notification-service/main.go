package main

import (
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func connectToRabbitMQ() *amqp.Connection {
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	if rabbitmqHost == "" {
		rabbitmqHost = "rabbitmq"
	}
	connStr := "amqp://sentinel:supersecretpassword@" + rabbitmqHost + ":5672/"

	for {
		conn, err := amqp.Dial(connStr)
		if err == nil {
			log.Println("Successfully connected to RabbitMQ.")
			return conn
		}
		log.Printf("Failed to connect to RabbitMQ: %s. Retrying in 5 seconds...", err)
		time.Sleep(5 * time.Second)
	}
}

func main() {
	conn := connectToRabbitMQ()
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"incidents.detected", // name
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received an incident: %s", d.Body)
			// Placeholder: Here you would add logic to send notifications
			// e.g., via email, Slack, Telegram, etc.
			log.Println("Processing notification...")
		}
	}()

	log.Printf("[*] Waiting for incidents. To exit press CTRL+C")
	<-forever
}
