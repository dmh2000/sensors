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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	user := os.Getenv("userid")
	pwd := os.Getenv("pwd")
	brk := os.Getenv("broker")

	// load configuration
	cfg, err := readConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	client, err := setupMQTT(cfg, user, pwd, brk)
	if err != nil {
		log.Println(err)
		os.Exit(5)
	}

	// add suscriptions
	err = subscribeMQTT(client, "w/sin")
	if err != nil {
		log.Println(err)
		os.Exit(3)
	}

	err = subscribeMQTT(client, "w/square")
	if err != nil {
		log.Println(err)
		os.Exit(3)
	}

	err = subscribeMQTT(client, "w/triangle")
	if err != nil {
		log.Println(err)
		os.Exit(3)
	}

	// block until user hits ctrl+c
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done // Will block here until user hits ctrl+c

	client.Disconnect(250)
}
