package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port          int    `json:"port"`
	DefaultTarget string `json:"default-target"`
}

func NewConfig() *Config {
	if _, err := os.Stat(ConfigPath); err != nil {
		log.Println("config not existing, creating default")
		data, err := json.Marshal(DefaultConfig)
		if err != nil {
			panic(err)
		}
		os.WriteFile(ConfigPath, data, 0660)
	}
	byt, err := os.ReadFile(ConfigPath)
	if err != nil {
		panic(err)
	}
	var dat Config
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	return &dat
}
