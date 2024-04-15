package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	// load configuration
	cfg, err := readConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// initialize the client
	var broker = brk // find the host name in the Overview of your cluster (see readme)
	var port = 8883  // find the port right under the host name, standard is 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("recv") // set a name as you desire
	opts.SetUsername(user)
	opts.SetPassword(pwd)

	// callback for subscribed messages
	opts.SetDefaultPublishHandler(messagePubHandlerFunc(cfg.debug))

	// callback when connected to broker
	opts.OnConnect = connectHandler

	// callback when connection is lost
	opts.OnConnectionLost = connectLostHandler

	// create the client using the options above
	client := mqtt.NewClient(opts)
	// throw an error if the connection isn't successfull
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(2)
	}

	// add suscriptions
	err = subscribe(client, "waveform/sin")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	err = subscribe(client, "waveform/square")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	err = subscribe(client, "waveform/triangle")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	// block until user hits ctrl+c
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done // Will block here until user hits ctrl+c

	client.Disconnect(250)
}
