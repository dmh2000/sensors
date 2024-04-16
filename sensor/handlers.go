package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// returns a callback function that will print the received message if 'debug' is true
func messagePubHandlerFunc(debug bool) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		if debug {
			fmt.Printf("%s,%s\n", msg.Topic(), msg.Payload())
		}
	}
}

// upon connection to the client, this is called
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	//fmt.Println("Connected")
}

// this is called when the connection to the client is lost, it prints "Connection lost" and the corresponding error
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	//fmt.Printf("Connection lost: %v", err)
}

// func subscribe(client mqtt.Client, cfg config) error {
// 	// subscribe to the same topic, that was published to, to receive the messages
// 	topic := fmt.Sprintf("waveform/%s", cfg.shape)
// 	token := client.Subscribe(topic, 1, nil)
// 	token.Wait()
// 	// Check for errors during subscribe (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
// 	if token.Error() != nil {
// 		return token.Error()
// 	}
// 	return nil
// }

// publish messages to the topic "topic/sensor"
// with a frequency in Hz and amplitude in m
func publish(client mqtt.Client, cfg config) error {
	// publish the message "Message" to the topic "topic/test" 10 times in a for loop
	ms := time.Millisecond * time.Duration(cfg.dt*1000)

	w := newWave(cfg.shape, cfg.amplitude, cfg.frequency, cfg.dt)
	for {
		// step the simulation
		x, y, _ := w.step()

		// format the message
		text := fmt.Sprintf("%.3f,%.3f", x, y)
		topic := fmt.Sprintf("waveform/%s", cfg.shape)
		token := client.Publish(topic, 0, false, text)
		token.Wait()
		// Check for errors during publishing (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
		if token.Error() != nil {
			return token.Error()
		}
		time.Sleep(ms)
	}
}
