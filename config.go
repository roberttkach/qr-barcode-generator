package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port string `json:"port"`
}

func readConfig() (Config, error) {
	var config Config
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(configFile, &config)
	return config, err
}
