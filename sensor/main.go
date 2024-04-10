package main

import (
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
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

	cfg, err := readConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var broker = brk // find the host name in the Overview of your cluster (see readme)
	var port = 8883  // find the port right under the host name, standard is 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("test") // set a name as you desire
	opts.SetUsername(user)
	opts.SetPassword(pwd)
	// (optionally) configure callback handlers that get called on certain events
	opts.SetDefaultPublishHandler(messagePubHandlerFunc(cfg.debug))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	// create the client using the options above
	client := mqtt.NewClient(opts)
	// throw an error if the connection isn't successfull
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(2)
	}

	err = subscribe(client, cfg.debug)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	err = publish(client, cfg.frequency, cfg.amplitude, cfg.dt)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	client.Disconnect(250)
}
