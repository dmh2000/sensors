package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// returns a callback function that will print the received message if 'debug' is true
func messagePubHandlerFunc(ch chan string) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		rbq := fmt.Sprintf("%s,%s", msg.Topic(), msg.Payload())
		// send to producer
		ch <- rbq
	}
}

// upon connection to the client, this is called
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

// this is called when the connection to the client is lost, it prints "Connection lost" and the corresponding error
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	//fmt.Printf("Connection lost: %v", err)
}

func subscribeMQTT(client mqtt.Client, topic string) error {
	// subscribe to the same topic, that was published to, to receive the messages
	fmt.Println("subscribing to topic: ", topic)
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	// Check for errors during subscribe (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
	if token.Error() != nil {
		return token.Error()
	}
	return nil
}

func setupMQTT(pipe chan string, user string, pwd string, brk string) (mqtt.Client, error) {
	// initialize the client
	var broker = brk // find the host name in the Overview of your cluster (see readme)
	var port = 8883  // find the port right under the host name, standard is 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("receiver") // set a name as you desire
	opts.SetUsername(user)
	opts.SetPassword(pwd)

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

	return client, nil
}

func subscriber(pipe chan string) {
	// setup mqtt receiver
	user := os.Getenv("userid")
	pwd := os.Getenv("pwd")
	brk := os.Getenv("broker")

	client, err := setupMQTT(pipe, user, pwd, brk)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(250)

	// add suscriptions
	err = subscribeMQTT(client, "w/sin")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	err = subscribeMQTT(client, "w/square")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	err = subscribeMQTT(client, "w/triangle")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	// block until user hits ctrl+c
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done // Will block here until user hits ctrl+c
}
