package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// returns a callback function that will print the received message if 'debug' is true
func messagePubHandlerFunc(debug bool) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		if debug {
			log.Printf("%s,%s\r", msg.Topic(), msg.Payload())
		}
	}
}

// upon connection to the client, this is called
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
}

// this is called when the connection to the client is lost, it prints "Connection lost" and the corresponding error
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("Connection lost: %v", err)
}

// publishMQTT messages to the topic "topic/sensor"
// with a frequency in Hz and amplitude in m
func publishMQTT(client mqtt.Client, cfg config) error {
	// publish the message "Message" to the topic "topic/test" 10 times in a for loop
	ms := time.Millisecond * time.Duration(cfg.dt*1000)
	ticker := time.NewTicker(ms)

	w, err := newWave(cfg.shape, cfg.amplitude, cfg.frequency, cfg.dt)
	if err != nil {
		return err
	}

	for {
		// step the simulation
		x, y := w.step()

		// format the message
		text := fmt.Sprintf("%.3f,%.3f", x, y)
		topic := fmt.Sprintf("w/%s", cfg.shape)
		token := client.Publish(topic, 0, false, text)
		token.Wait()
		// Check for errors during publishing the message
		if token.Error() != nil {
			return fmt.Errorf("mqtt publish: %w", token.Error())
		}

		<-ticker.C
	}
}

const retryDelay = 5

func setupMQTT(cfg config, user string, pwd string, url string) (mqtt.Client, error) {
	// initialize the client
	opts := mqtt.NewClientOptions()
	opts.AddBroker(url)
	opts.SetClientID(cfg.name) // set a name as you desire

	// set username and password if provided
	if user != "" && pwd != "" {
		opts.SetUsername(user)
		opts.SetPassword(pwd)
	}

	// callback for subscribed messages
	opts.SetDefaultPublishHandler(messagePubHandlerFunc(cfg.debug))

	// callback when connected to broker
	opts.OnConnect = connectHandler

	// callback when connection is lost
	opts.OnConnectionLost = connectLostHandler

	// create the client using the options above
	client := mqtt.NewClient(opts)
	for {
		// try to connect to the broker
		if token := client.Connect(); token.Wait() && token.Error() == nil {
			break
		}
		// if the connection fails, wait  before trying again
		// could use backoff here
		time.Sleep(retryDelay * time.Second)
	}
	log.Println("Connected to MQTT broker")

	return client, nil
}
