package main

import "time"

// Solicitation request for a bid.
type Solicitation struct {
	Title      string            `json:"title"`
	Properties map[string]string `json:"properties"`
	Modified   time.Time         `json:"modified"`
}
