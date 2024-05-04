package main

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func rabbitConsumer(c chan string) {
	var conn *amqp.Connection
	var err error
	for {
		// is it running in docker?
		conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err == nil {
			log.Println("API running in docker")
			break
		}
		log.Printf("API Not running in docker : %s\n", err)

		// is it running locally?
		conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err == nil {
			log.Println("API Running in localhost")
			break
		}
		log.Printf("API not running locally : %s\n", err)

		// wait and retry
		time.Sleep(5 * time.Second)
		log.Println("API reconnecting to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"waveforms", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
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

	for msg := range msgs {
		c <- string(msg.Body)
	}
}
