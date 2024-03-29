package main

import (
	"encoding/json"
	"log"
	"os"
)

type Shorts struct {
	Shorts []Short `json:"shorts"`
}

func NewShorts() *Shorts {
	if _, err := os.Stat("go-shorts.json"); err != nil {
		log.Println("shorts not existing, creating default")
		short := NewShort("git", "https://github.com/tim-krehan/go-link-shortener")
		var sl []Short
		sl = append(sl, *short)
		json_config := Shorts{
			Shorts: sl,
		}
		data, err := json.Marshal(json_config)
		if err != nil {
			panic(err)
		}
		os.WriteFile("go-shorts.json", data, 0660)
	}
	byt, err := os.ReadFile("go-shorts.json")
	if err != nil {
		panic(err)
	}
	var dat Shorts
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	return &dat
}

func (s *Shorts) AddShort(short Short) {
	s.Shorts = append(s.Shorts, short)
}

func (s Shorts) GetShort(slug string) *Short {
	for _, sh := range s.Shorts {
		if sh.Slug == slug {
			return &sh
		}
	}
	return nil
}
