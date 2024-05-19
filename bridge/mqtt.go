package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// receive msgs from mqtt and forward to the rabbitmq producer
func messagePubHandlerFunc(ch chan string) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		rbq := fmt.Sprintf("%s,%s", msg.Topic(), msg.Payload())

		select {
		case ch <- rbq:
			// log.Println(rbq)
		default:
			log.Println("Channel full, dropping message")
		}
	}
}

// upon connection to the client, this is called
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
}

// this is called when the connection to the client is lost, it prints "Connection lost" and the corresponding error
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Println(err)
}

const subscribeTimeout = 5

func subscribeMQTT(client mqtt.Client, topic string) error {
	// subscribe to the same topic, that was published to, to receive the messages
	token := client.Subscribe(topic, 1, nil)
	if token.Error() != nil {
		return token.Error()
	}
	token.WaitTimeout(subscribeTimeout * time.Second)
	if token.Error() == nil {
		log.Println("subscribed to topic: ", topic)
	}
	return token.Error()
}

func setupMQTT(pipe chan string, user string, pwd string, url string, id string) (mqtt.Client, error) {
	// initialize the client
	opts := mqtt.NewClientOptions()
	opts.AddBroker(url)
	opts.SetClientID(id) // generate a unique client ID using `github.com/google/uuid

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

const defaultUrl = "tcp://localhost:1883"

func subscriber(pipe chan string, topics []string) {
	// setup mqtt receiver
	user, _ := os.LookupEnv("userid") // default to empty string
	pwd, _ := os.LookupEnv("pwd")     // default to empty string
	url, ok := os.LookupEnv("url")    // use default url
	if !ok {
		url = defaultUrl
	}

	client, err := setupMQTT(pipe, user, pwd, url, "bridge")
	if err != nil {
		// kill the app if parameters are incorrect
		log.Fatal(err)
	}
	defer client.Disconnect(250)

	for _, topic := range topics {
		// add suscriptions
		err = subscribeMQTT(client, topic)
		if err != nil {
			log.Println(err)
		}
	}

	// block until user hits ctrl+c
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done // Will block here until user hits ctrl+c
}
