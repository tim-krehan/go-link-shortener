package main

import (
	"encoding/json"
	"log"
	"net/url"
)

type Short struct {
	Slug      string
	Target    string
	targetUrl url.URL
}

func NewShort(slug, target string) *Short {
	u, _ := url.Parse(target)
	s := Short{
		Slug:      slug,
		Target:    target,
		targetUrl: *u,
	}
	log.Printf("loaded \"%v\" to \"%v\"", s.Slug, s.Target)
	return &s
}
func (s *Short) LoadShort(slug, target string) {
	u, _ := url.Parse(target)
	s.Slug = slug
	s.Target = target
	s.targetUrl = *u
	log.Printf("loaded \"%v\" to \"%v\"", s.Slug, s.Target)
}

func (short *Short) UnmarshalJSON(b []byte) error {
	var s map[string]string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	for key, value := range s {
		short.LoadShort(key, value)
	}
	return nil
}
func (short Short) MarshalJSON() ([]byte, error) {
	slug := map[string]string{
		short.Slug: short.Target,
	}
	return json.Marshal(slug)
}
