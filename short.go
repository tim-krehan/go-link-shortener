package main

import (
	"log"
	"net/url"
)

type Short struct {
	Slug        string `json:"slug"`
	Target      string `json:"target"`
	targetUrl   url.URL
	Description string `json:"description"`
}

func NewShort(slug, target, description string) *Short {
	u, _ := url.Parse(target)
	s := Short{
		Slug:        slug,
		Target:      target,
		targetUrl:   *u,
		Description: description,
	}
	log.Printf("loaded \"%v\" => \"%v\" - %v", s.Slug, s.Target, s.Description)
	return &s
}
func (s *Short) LoadShort() {
	u, _ := url.Parse(s.Target)
	s.targetUrl = *u
	log.Printf("loaded \"%v\" => \"%v\" - %v", s.Slug, s.Target, s.Description)
}
