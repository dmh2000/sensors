package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	// load secrets
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// pipe between mqtt receiver and rabbitmq producer
	pipe := make(chan string, 10)

	// subscribe to mqtt topics
	go subscriber(pipe)

	// publish to rabbitmq
	go producer(pipe)

	// block until user hits ctrl+c
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done // Will block here until user hits ctrl+c

}
