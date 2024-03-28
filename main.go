package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Short struct {
	Slug      string `json:"slug"`
	targetUrl url.URL
	Target    string `json:"target"`
}
type Config struct {
	Port   int     `json:"port"`
	Shorts []Short `json:"shorts"`
}

func NewRedir(slug string, target string) *Short {
	t, err := url.Parse(target)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	r := Short{
		Slug:      slug,
		Target:    target,
		targetUrl: *t,
	}
	log.Printf("loaded \"%v\" to \"%v\"", r.Slug, r.targetUrl.String())
	return &r
}

func redirect(w http.ResponseWriter, req *http.Request) {
	if len(strings.Split(req.URL.Path, "/")) >= 2 {
		target := strings.Split(req.URL.Path, "/")[2]
		log.Printf("selected target is \"%v\"\n", target)
	} else {
		panic("no target!")
	}
}

func load_config() *Config {
	if _, err := os.Stat("go-link-shortener.cfg"); err != nil {
		log.Println("config not existing, creating default")
		json_config := Config{
			Port:   8080,
			Shorts: []Short{*NewRedir("git", "https://github.com/tim-krehan/go-link-shortener")},
		}
		data, err := json.Marshal(json_config)
		if err != nil {
			panic(err)
		}
		os.WriteFile("go-link-shortener.cfg", data, 0660)
	}
	byt, err := os.ReadFile("go-link-shortener.cfg")
	if err != nil {
		panic(err)
	}
	var dat Config
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	return &dat
}

func main() {
	log.Println("initializing service")
	config := load_config()
	log.Println("loaded config")
	log.Printf("%+v", config)

	log.Println("registering route \"/to/*\"")
	http.HandleFunc("/to/*", redirect)

	log.Printf("listening to new requests on :%v\n", config.Port)
	http.ListenAndServe(fmt.Sprint(":", config.Port), nil)
}
