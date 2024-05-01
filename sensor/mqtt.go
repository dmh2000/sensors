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
			fmt.Printf("%s,%s\r", msg.Topic(), msg.Payload())
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

// publishMQTT messages to the topic "topic/sensor"
// with a frequency in Hz and amplitude in m
func publishMQTT(client mqtt.Client, cfg config) error {
	// publish the message "Message" to the topic "topic/test" 10 times in a for loop
	ms := time.Millisecond * time.Duration(cfg.dt*1000)

	w := newWave(cfg.shape, cfg.amplitude, cfg.frequency, cfg.dt)
	for {
		// step the simulation
		x, y := w.step()

		// format the message
		text := fmt.Sprintf("%.3f,%.3f", x, y)
		topic := fmt.Sprintf("w/%s", cfg.shape)
		token := client.Publish(topic, 0, false, text)
		token.Wait()
		// Check for errors during publishing (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
		if token.Error() != nil {
			return token.Error()
		}
		time.Sleep(ms)
	}
}

func setupMQTT(cfg config, user string, pwd string, brk string) (mqtt.Client, error) {
	// initialize the client
	var broker = brk // find the host name in the Overview of your cluster (see readme)
	var port = 8883  // find the port right under the host name, standard is 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID(cfg.name) // set a name as you desire
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
		return nil, token.Error()
	}

	return client, nil
}
