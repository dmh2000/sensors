package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	shape     string
	frequency float64 // simulated wave frequency
	amplitude float64 // simulated wave amplitude
	dt        float64 // computed sample time dt
	debug     bool    // debug flag
}

func readConfig() (config, error) {
	var cfg config

	// load configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return cfg, fmt.Errorf("error loading config file")
	}
	shape := viper.Get("shape")
	frequency := viper.Get("frequency")
	amplitude := viper.Get("amplitude")
	rate := viper.Get("rate")
	debug := viper.Get("debug")

	if shape == nil || frequency == nil || amplitude == nil || rate == nil {
		return cfg, fmt.Errorf("missing configuration parameters")
	}

	w, ok := shape.(string)
	if !ok {
		return cfg, fmt.Errorf("config 'shape' is not a string")
	}

	f, ok := frequency.(float64)
	if !ok {
		return cfg, fmt.Errorf("config 'frequency' is not a float64")
	}
	a, ok := amplitude.(float64)
	if !ok {
		return cfg, fmt.Errorf("config 'amplitude' is not a float64")
	}
	r, ok := rate.(float64)
	if !ok {
		return cfg, fmt.Errorf("config 'rate' is not a float")
	}

	p, ok := debug.(bool)
	if !ok {
		return cfg, fmt.Errorf("config 'debug' is not a bool")
	}

	// computer delta T in fractional seconds
	dt := 1.0 / float64(r)

	cfg = config{w, f, a, dt, p}

	return cfg, nil
}
