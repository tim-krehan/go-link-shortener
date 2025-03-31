package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fsnotify/fsnotify"
)

var ConfigPath = "go-link-shortener.json"
var ShortsPath = "go-shorts.json"

var AllShorts *Shorts

var DefaultConfig = Config{
	Port:          8081,
	DefaultTarget: "https://github.com/tim-krehan/go-link-shortener",
}

func main() {
	log.Println("initializing service")
	log.Println("load config")
	config := NewConfig()
	log.Println("loaded config")

	log.Println("load shorts")
	AllShorts, _ = NewShorts()
	log.Println("loaded shorts")

	log.Println("registering route \"/to/*\"")
	http.HandleFunc("/to/", redirect)

	log.Println("registering route \"/list\"")
	http.HandleFunc("/list", list)

	log.Println("creating monitor for config changes")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	watcher.Add(".")
	go AllShorts.MonitorConfig(watcher)

	log.Printf("listening to new requests on :%v\n", config.Port)
	go http.ListenAndServe(fmt.Sprint(":", config.Port), nil)

	<-make(chan interface{})
}
