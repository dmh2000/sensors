package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	err := rabbitConsumer()
	if err != nil {
		log.Fatalf("Error RabbitMQ setup: %s", err)
	}

	// block until user hits ctrl+c
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done // Will block here until user hits ctrl+c
}
