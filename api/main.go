package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// record latest update
var sinX string = "0.0"
var sinY string = "0.0"
var squareX string = "0.0"
var squareY string = "0.0"
var triangleX string = "0.0"
var triangleY string = "0.0"

func main() {
	data := make(chan string)

	// connect to rabbitMQ and start listening for messages
	go func() {
		err := rabbitConsumer(data)
		if err != nil {
			log.Fatalf("Error connecting to RabbitMQ: %s", err)
		}
	}()

	// read from the channel and update the latest values
	go func() {
		for d := range data {
			s := strings.Split(d, ",")
			switch s[0] {
			case "w/sin":
				sinX = s[1]
				sinY = s[2]
			case "w/square":
				squareX = s[1]
				squareY = s[2]
			case "w/triangle":
				triangleX = s[1]
				triangleY = s[2]
			}
			// log.Printf("%s,%s,%s,%s,%s,%s\n", sinX, sinY, squareX, squareY, triangleX, triangleY)
		}
	}()

	// create a new router
	router := http.NewServeMux()

	router.HandleFunc("GET /sensor/v1/sin", func(w http.ResponseWriter, r *http.Request) {
		// return latest sin wave
		fmt.Fprintf(w, "sin,%s,%s", sinX, sinY)
	})

	router.HandleFunc("GET /sensor/v1/square", func(w http.ResponseWriter, r *http.Request) {
		// return latest square wave
		fmt.Fprintf(w, "square,%s,%s", squareX, squareY)
	})

	router.HandleFunc("GET /sensor/v1/triangle", func(w http.ResponseWriter, r *http.Request) {
		// return latest triangle wave
		fmt.Fprintf(w, "triangle,%s,%s", triangleX, triangleY)
	})

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the sensor API")
	})

	http.ListenAndServe(":8002", router)
}
