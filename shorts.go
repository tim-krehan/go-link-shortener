package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)

type Shorts struct {
	Shorts []Short `json:"shorts"`
}

func NewShorts() (*Shorts, error) {
	var dat Shorts
	var err error
	if _, err := os.Stat(ShortsPath); err != nil {
		log.Println("shorts not existing, creating default")
		short := NewShort("git", "https://github.com/tim-krehan/go-link-shortener", "Link shortener written in go. even includes a list with all shorts :).")
		var sl []Short
		sl = append(sl, *short)
		json_config := Shorts{
			Shorts: sl,
		}
		data, err := json.Marshal(json_config)
		if err != nil {
			log.Println("ERROR failed create default shorts:", err)
			return nil, err
		}
		os.WriteFile(ShortsPath, data, 0660)
	}
	log.Println("reading shorts config from file")
	byt, err := os.ReadFile(ShortsPath)
	if err != nil {
		log.Println("ERROR failed to read file:", err)
		return nil, err
	}
	log.Println("parsing config")
	err = json.Unmarshal(byt, &dat)
	for index := range dat.Shorts {
		dat.Shorts[index].LoadShort()
	}

	if err != nil {
		log.Println("ERROR failed to load shorts:", err)
		log.Println("using default shorts")
		short := NewShort("git", "https://github.com/tim-krehan/go-link-shortener", "Link shortener written in go. even includes a list with all shorts :).")
		dat.Shorts = append(dat.Shorts, *short)
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
	var timer *time.Timer
	var last_event fsnotify.Event
	var nil_event fsnotify.Event
	timer = time.NewTimer(1 * time.Second)
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op == fsnotify.Write && event.Name == ("."+string(os.PathSeparator)+ShortsPath) {
				timer = time.NewTimer(1 * time.Second)
				last_event = event
				log.Println("registered event", event.Op, "for", event.Name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Error", err)
		case <-timer.C:
			if last_event != nil_event {
				log.Println("timer went off")
				newShorts, err := NewShorts()
				if err == nil {
					s.Shorts = newShorts.Shorts
				}
			}
		}
	}
}
