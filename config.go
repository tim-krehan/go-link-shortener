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
	if _, err := os.Stat("go-link-shortener.json"); err != nil {
		log.Println("config not existing, creating default")
		data, err := json.Marshal(DefaultConfig)
		if err != nil {
			panic(err)
		}
		os.WriteFile("go-link-shortener.json", data, 0660)
	}
	byt, err := os.ReadFile("go-link-shortener.json")
	if err != nil {
		panic(err)
	}
	var dat Config
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	return &dat
}
