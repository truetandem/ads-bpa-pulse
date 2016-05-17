package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/appengine/datastore"

	"golang.org/x/net/context"
)

// Solicitation request for a bid.
type Solicitation struct {
	Title      string            `json:"title" datastore:"Title"`
	Properties map[string]string `json:"properties" datastore:"-"`
	Modified   time.Time         `json:"modified" datastore:"Modified,noindex"`
	checksum   string            `datastore:"Checksum"`
	properties []byte            `datastore:"Properties"`
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

// Get a Solicitation from the datastore.
func (s *Solicitation) Get(ctx context.Context) (Solicitation, error) {
	// Create a new key based on the title and attempt get the
	// entity
	var sol Solicitation
	key := datastore.NewKey(ctx, "Solicitation", s.Title, 0, nil)
	if err := datastore.Get(ctx, key, sol); err != nil {
		return sol, err
	}

	// Unmarshal the binary JSON and store it as a map
	if err := json.Unmarshal(sol.properties, sol.Properties); err != nil {
		return sol, err
	}

	return sol, nil
}

// Save a Solicitation to the datastore.
func (s *Solicitation) Save(ctx context.Context) error {
	// Save the checksum to an internal variable for storage
	// purposes
	s.checksum = s.Checksum()

	// Marshal the property map so we can save it as a binary
	// field
	js, err := json.Marshal(s.Properties)
	if err != nil {
		return err
	}
	s.properties = js

	// Create the key based on the title and store it
	key := datastore.NewKey(ctx, "Solicitation", s.Title, 0, nil)
	if _, err := datastore.Put(ctx, key, s); err != nil {
		return err
	}

	return nil
}
