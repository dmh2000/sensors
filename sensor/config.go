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
	name      string  // client name
}

var validShape = map[string]bool{
	"square":   true,
	"sin":      true,
	"triangle": true,
}

// readConfig reads the configuration from the given file and returns a
// config struct or an error if the configuration is missing or invalid.
// configFile: name of the configuration file
// Returns: config struct or error
func readConfig(configFile string) (config, error) {
	var cfg config

	// load configuration
	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return cfg, fmt.Errorf("error loading config file")
	}

	// get configuration parameters
	shape := viper.Get("shape")
	frequency := viper.Get("frequency")
	amplitude := viper.Get("amplitude")
	rate := viper.Get("rate")
	debug := viper.Get("debug")
	name := viper.Get("name")

	// check if all parameters are present
	if shape == nil || frequency == nil || amplitude == nil || rate == nil {
		return cfg, fmt.Errorf("missing configuration parameters")
	}

	// check if all parameters are of the correct type
	w, ok := shape.(string)
	if !ok {
		return cfg, fmt.Errorf("config 'shape' is not a string")
	}
	if !validShape[w] {
		return cfg, fmt.Errorf("unknown shape: %s", w)
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

	n, ok := name.(string)
	if !ok {
		return cfg, fmt.Errorf("config 'name' is not a string")
	}

	// compute delta T in fractional seconds
	dt := 1.0 / r

	// create config struct
	cfg = config{w, f, a, dt, p, n}

	return cfg, nil
}
