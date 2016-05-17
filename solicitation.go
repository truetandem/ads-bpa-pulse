package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

// Solicitation request for a bid.
type Solicitation struct {
	Title      string            `json:"title"`
	Properties map[string]string `json:"properties"`
	Modified   time.Time         `json:"modified"`
}

// ToString returns the Solicitation as a formatted string.
func (s *Solicitation) ToString() string {
	str := fmt.Sprintf("Title: %s\n", s.Title)
	for k, v := range s.Properties {
		str += fmt.Sprintf("\n  - %s: %s", k, v)
	}

	return str
}

// Checksum returns the SHA-1 checksum of the Solicitation.
func (s *Solicitation) Checksum() string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s.ToString())))
}
