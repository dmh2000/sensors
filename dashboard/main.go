package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

type Sensor struct {
	X float64
	Y float64
}
type Sensors []Sensor

type Wave struct {
	W string
	X string
	Y string
}

type Waves struct {
	Sin Wave
	Sqr Wave
	Tri Wave
}

func fetchApiWave(url string) (Wave, error) {
	var w Wave
	// fetch the api data
	api, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching %s wave data: %s\n", url, err)
		return w, err
	}
	defer api.Body.Close()

	body, err := io.ReadAll(api.Body)
	if err != nil {
		log.Printf("Error reading %s wave data: %s\n", url, err)
		return w, err
	}

	s := string(body)
	data := strings.Split(s, ",")

	w.W = data[0]
	w.X = data[1]
	w.Y = data[2]

	return w, nil
}

func main() {

	update, err := template.ParseFiles("update.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	index, err := template.ParseFiles("index.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	var waves Waves

	// create a new router
	router := http.NewServeMux()

	router.HandleFunc("GET /update", func(w http.ResponseWriter, r *http.Request) {

		sin, err := fetchApiWave("http://api:8002/sensor/v1/sin")
		if err != nil {
			log.Printf("Error fetching sin wave data: %s\n", err)
		}
		// log.Printf("sin: %v\n", sin)
		waves.Sin.X = sin.X
		waves.Sin.Y = sin.Y
		waves.Sin.W = sin.W

		sqr, err := fetchApiWave("http://api:8002/sensor/v1/square")
		if err != nil {
			log.Printf("Error fetching square wave data: %s\n", err)
		}
		// log.Printf("sqr: %v\n", sqr)
		waves.Sqr.X = sqr.X
		waves.Sqr.Y = sqr.Y
		waves.Sqr.W = sqr.W

		tri, err := fetchApiWave("http://api:8002/sensor/v1/triangle")
		if err != nil {
			log.Printf("Error fetching triangle wave data: %s\n", err)
		}
		// log.Printf("tri: %v\n", tri)
		waves.Tri.X = tri.X
		waves.Tri.Y = tri.Y
		waves.Tri.W = tri.W

		err = update.Execute(w, waves)
		if err != nil {
			log.Fatal(err)
		}
	})

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {

		err := index.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	http.ListenAndServe(":8001", router)
}
