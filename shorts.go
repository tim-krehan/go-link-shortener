package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

type Shorts struct {
	Shorts []Short `json:"shorts"`
}

func NewShorts() (*Shorts, error) {
	if _, err := os.Stat(ShortsPath); err != nil {
		log.Println("shorts not existing, creating default")
		short := NewShort("git", "https://github.com/tim-krehan/go-link-shortener")
		var sl []Short
		sl = append(sl, *short)
		json_config := Shorts{
			Shorts: sl,
		}
		data, err := json.Marshal(json_config)
		if err != nil {
			return nil, err
		}
		os.WriteFile(ShortsPath, data, 0660)
	}
	byt, err := os.ReadFile(ShortsPath)
	if err != nil {
		return nil, err
	}
	var dat Shorts
	if err := json.Unmarshal(byt, &dat); err != nil {
		return nil, err
	}
	return &dat, err
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

func (s *Shorts) MonitorConfig(watcher *fsnotify.Watcher) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op == fsnotify.Write {
				log.Println("Write Event registered to", event.Name)
				log.Println("reloading config")
				newShorts, err := NewShorts()
				if err != nil {
					log.Println("failed to load config:", err)
				}
				s = newShorts
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Error", err)
		}
	}
}
