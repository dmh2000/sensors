package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func producer(pipe chan string) {
	var conn *amqp.Connection
	var err error
	// connect to RabbitMQ
	for {
		// is it running in docker?
		conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err == nil {
			log.Println("Bridge running in docker")
			break
		}
		log.Printf("Bridge Not running in docker : %s\n", err)

		// is it running locally?
		conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err == nil {
			log.Println("Bridge Running in localhost")
			break
		}
		log.Printf("Bridge not running locally : %s\n", err)

		// wait and retry
		time.Sleep(5 * time.Second)
		log.Println("API reconnecting to RabbitMQ")
	}
	defer conn.Close()

	// open a channel
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// receive messages from MQTT and publish to RabbitMQ
	for msg := range pipe {
		body := msg
		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")
	}

	log.Fatal("pipe closed")
}
