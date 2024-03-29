package main

import (
	"fmt"
	"log"
	"net/http"
)

var AllShorts *Shorts

var DefaultConfig = Config{
	Port:          8080,
	DefaultTarget: "https://github.com/tim-krehan/go-link-shortener",
}

func main() {
	log.Println("initializing service")
	log.Println("load config")
	config := NewConfig()
	log.Println("loaded config")
	// log.Printf("%+v", config)

	log.Println("load shorts")
	AllShorts = NewShorts()
	log.Println("loaded shorts")

	log.Println("registering route \"/to/*\"")
	http.HandleFunc("/to/*", redirect)

	log.Printf("listening to new requests on :%v\n", config.Port)
	http.ListenAndServe(fmt.Sprint(":", config.Port), nil)
}
