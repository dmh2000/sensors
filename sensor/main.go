package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	configFile := "config"
	if len(os.Args) >= 2 {
		configFile = os.Args[1]
	}

	// load secrets
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	user := os.Getenv("userid")
	pwd := os.Getenv("pwd")
	url := os.Getenv("url")

	// load configuration
	cfg, err := readConfig(configFile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Printf("SENSOR,%s,%s.yaml,%v\n", url, configFile, cfg)

	log.Println(cfg, user, pwd, url)
	client, err := setupMQTT(cfg, user, pwd, url, "")
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
