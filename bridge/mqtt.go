package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// receive msgs from mqtt and forward to the rabbitmq producer
func messagePubHandlerFunc(ch chan string) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		rbq := fmt.Sprintf("%s,%s", msg.Topic(), msg.Payload())
		// send to producer
		ch <- rbq
	}
}

// upon connection to the client, this is called
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
}

// this is called when the connection to the client is lost, it prints "Connection lost" and the corresponding error
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	//log.Printf("Connection lost: %v", err)
}

func subscribeMQTT(client mqtt.Client, topic string) error {
	// subscribe to the same topic, that was published to, to receive the messages
	log.Println("subscribing to topic: ", topic)
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	// Check for errors during subscribe (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
	if token.Error() != nil {
		return token.Error()
	}
	return nil
}

func setupMQTT(pipe chan string, user string, pwd string, url string, id string) (mqtt.Client, error) {
	// initialize the client
	opts := mqtt.NewClientOptions()
	opts.AddBroker(url)
	opts.SetClientID(id) // set a name as you desire
	opts.SetUsername(user)
	opts.SetPassword(pwd)

	opts.AddBroker(url)
	opts.SetClientID(id) // set a name as you desire

	if user != "" && pwd != "" {
		opts.SetUsername(user)
		opts.SetPassword(pwd)
	}

	// callback for subscribed messages
	opts.SetDefaultPublishHandler(messagePubHandlerFunc(pipe))

	// callback when connected to broker
	opts.OnConnect = connectHandler

	// callback when connection is lost
	opts.OnConnectionLost = connectLostHandler

	// create the client using the options above
	client := mqtt.NewClient(opts)
	// throw an error if the connection isn't successfull
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	log.Println("Connected to MQTT broker")
	return client, nil
}

func subscriber(pipe chan string) {
	// setup mqtt receiver
	user := os.Getenv("userid")
	pwd := os.Getenv("pwd")
	url := os.Getenv("url")

	client, err := setupMQTT(pipe, user, pwd, url, "bridge")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(250)

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
}
