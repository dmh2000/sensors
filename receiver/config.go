package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	debug bool   // debug flag
	name  string // client name
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
	debug := viper.Get("debug")
	name := viper.Get("name")

	if debug == nil || name == nil {
		return cfg, fmt.Errorf("missing configuration parameters")
	}

	d, ok := debug.(bool)
	if !ok {
		return cfg, fmt.Errorf("config 'debug' is not a bool")
	}

	n, ok := name.(string)
	if !ok {
		return cfg, fmt.Errorf("config 'name' is not a string")
	}

	cfg = config{d, n}

	return cfg, nil
}
