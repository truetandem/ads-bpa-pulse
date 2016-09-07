package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"google.golang.org/appengine/datastore"

	"golang.org/x/net/context"
)

// Solicitation request for a bid.
type Solicitation struct {
	Title      string            `json:"title" datastore:"Title"`
	Properties map[string]string `json:"properties" datastore:"-"`
	Modified   time.Time         `json:"modified" datastore:"Modified,noindex"`
	Checksum   string            `datastore:"Checksum"`
	JSON       []byte            `datastore:"Json"`
}

// String returns the Solicitation as a formatted string.
func (s *Solicitation) String() string {
	str := fmt.Sprintf("Title: %s\n", s.Title)

	// Sort the properties because ranging over a map
	// returns elements in a random order
	mk := make([]string, len(s.Properties))
	i := 0
	for k := range s.Properties {
		mk[i] = k
		i++
	}
	sort.Strings(mk)

	for _, k := range mk {
		str += fmt.Sprintf("\n  - %s: %s", k, s.Properties[k])
	}

	return str
}

// Sum returns the SHA-1 checksum of the Solicitation.
func (s *Solicitation) Sum() string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s.String())))
}

// Get a Solicitation from the datastore.
func (s *Solicitation) Get(ctx context.Context) (Solicitation, error) {
	// Create a new solicitation
	sol := Solicitation{
		Properties: map[string]string{},
	}

	// Search by the title and attempt get the entity
	key := datastore.NewKey(ctx, "Solicitation", s.Title, 0, nil)
	if err := datastore.Get(ctx, key, &sol); err != nil {
		return sol, err
	}

	// Unmarshal the binary JSON and store it as a map
	if err := json.Unmarshal(sol.JSON, &sol.Properties); err != nil {
		return sol, err
	}

	return sol, nil
}

// Save a Solicitation to the datastore.
func (s *Solicitation) Save(ctx context.Context) error {
	// Save the checksum to an internal variable for storage
	// purposes
	s.Checksum = s.Sum()

	// Marshal the property map so we can save it as a binary
	// field
	js, err := json.Marshal(s.Properties)
	if err != nil {
		return err
	}
	s.JSON = js

	// Create the key based on the title and store it
	key := datastore.NewKey(ctx, "Solicitation", s.Title, 0, nil)
	if _, err := datastore.Put(ctx, key, s); err != nil {
		return err
	}

	return nil
}
