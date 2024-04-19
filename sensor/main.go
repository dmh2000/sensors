package main

import (
	"fmt"
	"log"
	"os"

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
		fmt.Println(err)
		os.Exit(1)
	}

	client, err := setupMQTT(cfg, user, pwd, brk)
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	// add publications, does not return unless there is an error
	err = publishMQTT(client, cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	client.Disconnect(250)
}
