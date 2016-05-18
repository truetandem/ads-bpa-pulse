package main

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"google.golang.org/appengine/datastore"

	"golang.org/x/net/context"
)

var (
	ErrEmailRequired      = errors.New("Email is required to create new subscription")
	ErrEmailInvalid       = errors.New("Email provided is invalid")
	ErrSubscriptionExists = errors.New("Email address already subscribed")
	EmailRegexp           *regexp.Regexp
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

// Retrieves a subscription using a users email address. Struct pointer is passed in
// so current object gets populated with data
func (s *Subscription) Get(ctx context.Context) (*datastore.Key, error) {
	key := datastore.NewKey(ctx, "Subscription", s.Email, 0, nil)
	return key, datastore.Get(ctx, key, s)
}

// Creates a new subscription for a particular email address. Ensures that a subscription
// does not exist for the given email.
func (s *Subscription) Subscribe(ctx context.Context) (*datastore.Key, error) {
	// Attempt to retrieve subscription to see if one exists
	if _, err := s.Get(ctx); err != nil {
		// No entity found, so let's save!
		if err == datastore.ErrNoSuchEntity {
			return s.Save(ctx)
		}

		// Something else bad happened
		return nil, err
	}

	// We have a datastore hit.
	return nil, ErrSubscriptionExists
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

func (s *Subscription) String() string {
	return fmt.Sprintf("Subscription - Email:  [%v] Modfied: [%v]", s.Email, s.Modified)
}
