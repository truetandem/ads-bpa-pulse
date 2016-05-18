package main

import (
	"errors"
	"regexp"
	"time"

	"google.golang.org/appengine/datastore"

	"golang.org/x/net/context"
)

var (
	ErrEmailRequired = errors.New("Email is required to create new subscription")
	ErrEmailInvalid  = errors.New("Email provided is invalid")
	EmailRegexp      *regexp.Regexp
)

func init() {
	EmailRegexp = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
}

type Subscription struct {
	Email    string
	Modified time.Time
}

// Saves a new subscription. Makes sure that a valid email is passed in
func (s *Subscription) Save(ctx context.Context) (*datastore.Key, error) {
	// Make sure user provides a valid email address
	if valid, err := s.ValidEmail(); !valid {
		return nil, err
	}

	// Set timestamp
	s.Modified = time.Now()

	// Generate key using email address as string id
	key := datastore.NewKey(ctx, "Subscription", s.Email, 0, nil)
	_, err := datastore.Put(ctx, key, s)
	return key, err
}

// Ensures that email provided is valid
func (s *Subscription) ValidEmail() (bool, error) {
	if s.Email == "" {
		return false, ErrEmailRequired
	}

	if !EmailRegexp.MatchString(s.Email) {
		return false, ErrEmailInvalid
	}
	return true, nil
}
