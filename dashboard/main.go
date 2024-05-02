package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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

func main() {

	update, err := template.ParseFiles("update.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	index, err := template.ParseFiles("index.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	sensors := Sensors{
		{X: 1.0, Y: 2.0},
		{X: 3.0, Y: 4.0},
		{X: 5.0, Y: 6.0},
	}

	// create a new router
	router := http.NewServeMux()

	router.HandleFunc("GET /update", func(w http.ResponseWriter, r *http.Request) {

		sensors[0].X += 0.1

		waves := Waves{
			// populate the Wave struct
			Sin: Wave{
				W: "sin",
				X: fmt.Sprintf("%6.3f", sensors[0].X),
				Y: fmt.Sprintf("%6.3f", sensors[0].Y),
			},
			Sqr: Wave{
				W: "sqr",
				X: fmt.Sprintf("%6.3f", sensors[1].X),
				Y: fmt.Sprintf("%6.3f", sensors[1].Y),
			},
			Tri: Wave{
				W: "tri",
				X: fmt.Sprintf("%6.3f", sensors[2].X),
				Y: fmt.Sprintf("%6.3f", sensors[2].Y),
			},
		}

		err := update.Execute(w, waves)
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

	http.ListenAndServe("localhost:8000", router)
}
