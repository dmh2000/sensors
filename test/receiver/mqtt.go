package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const mqttQOS = 2 // exactly once

// returns a callback function that will print the received message if 'debug' is true
func messagePubHandlerFunc(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("%s\n", msg.Payload())
}

// upon connection to the client, this is called
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	// log.Println("Connected")
}

// this is called when the connection to the client is lost, it prints "Connection lost" and the corresponding error
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	//log.Printf("Connection lost: %v", err)
}

func subscribeMQTT(client mqtt.Client, topic string) error {
	// subscribe to the same topic, that was published to, to receive the messages
	token := client.Subscribe(topic, mqttQOS, nil)
	token.Wait()
	// Check for errors during subscribe (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
	if token.Error() != nil {
		return token.Error()
	}
	return nil
}

func setupMQTT(cfg config, user string, pwd string, url string) (mqtt.Client, error) {
	// initialize the client
	opts := mqtt.NewClientOptions()
	opts.AddBroker(url)
	opts.SetClientID(cfg.name) // set a name as you desire
	opts.SetUsername(user)
	opts.SetPassword(pwd)

	// callback for subscribed messages
	opts.SetDefaultPublishHandler(messagePubHandlerFunc)

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

	return client, nil
}
