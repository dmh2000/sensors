package main

import (
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
	url := os.Getenv("url")

	// load configuration
	cfg, err := readConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println(cfg, user, pwd, url)
	client, err := setupMQTT(cfg, user, pwd, url)
	if err != nil {
		log.Println(err)
		os.Exit(5)
	}

	// add publications, does not return unless there is an error
	err = publishMQTT(client, cfg)
	if err != nil {
		log.Println(err)
		os.Exit(4)
	}

	client.Disconnect(250)
}
