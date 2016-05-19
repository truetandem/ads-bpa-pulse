package main

import (
	"testing"

	"google.golang.org/appengine/aetest"
)

func TestSubscriptionSave(t *testing.T) {
	ctx, done, _ := aetest.NewContext()
	defer done()

	email := "johnsflores@gmail.com"
	s := Subscription{Email: email}
	key, err := s.Save(ctx)

	if err != nil {
		t.Fatalf("Unable to Save Subscription [%v]\n", err)
	}

	if key == nil {
		t.Fatal("Key was not set")
	}

	if key.StringID() != email {
		t.Fatal("Key string id does not match email")
	}
}

func TestSubscriptionGet(t *testing.T) {
	ctx, done, _ := aetest.NewContext()
	defer done()

	email := "johnsflores@gmail.com"
	s := Subscription{Email: email}
	if _, err := s.Save(ctx); err != nil {
		t.Fatalf("Unable to Save Subscrition [%v]", err)
	}

	s2 := Subscription{Email: email}
	if _, err := s2.Get(ctx); err != nil {
		t.Fatalf("Unable to retrieve subscription using email [%v]", s2.Email)
	}

	if s.Email != s2.Email {
		t.Errorf("Subscription email mismatch. Expected [%v] Got [%v]", s.Email, s2.Email)
	}

	if s.Modified.Unix() != s2.Modified.Unix() {
		t.Errorf("Subscription Modified mismatch. Expected [%v] Got [%v]", s.Modified, s2.Modified)
	}
}

func TestSubscriptionSubscribe(t *testing.T) {
	ctx, done, _ := aetest.NewContext()
	defer done()
	email := "winston@flores.org"

	s := Subscription{Email: email}

	if _, err := s.Subscribe(ctx); err != nil {
		t.Fatalf("Could not subscribe user with email [%v] Err [%v]", s.Email, err)
	}

	s2 := Subscription{Email: email}
	if _, err := s2.Subscribe(ctx); err == nil {
		t.Fatalf("Subscribed multiple times using same email address. This should not happen")
	} else {
		if err != ErrSubscriptionExists {
			t.Fatalf("Attempted to subscribe multiple times. Expected Err [%v] Got Err [%v]", ErrSubscriptionExists, err)
		}
	}
}

func TestSubscriptionUnsubscribe(t *testing.T) {
	ctx, done, _ := aetest.NewContext()
	defer done()
	email := "winston@flores.org"

	s := Subscription{Email: email}

	if _, err := s.Subscribe(ctx); err != nil {
		t.Fatalf("Could not subscribe user with email [%v] Err [%v]", s.Email, err)
	}

	s2 := Subscription{Email: email}
	if err := s2.Unsubscribe(ctx); err != nil {
		t.Fatalf("Could not unsubscribe user with email [%v] Err [%v]", s2.Email, err)
	}

}

func TestSubscriptionValidEmail(t *testing.T) {
	var emailTests = []struct {
		email    string
		expected bool
	}{
		{"johnsflores@gmail.com", true},
		{"john.s.flores@gmail.com", true},
		{"john.s.flores@somewhere.gmail.com", true},
		{"johnsfloresgmail.com", false},
		{"johnsflores@", false},
		{"johnsflores@blah", false},
		{"", false},
	}

	for _, test := range emailTests {
		s := Subscription{Email: test.email}
		if valid, err := s.ValidEmail(); valid != test.expected {
			t.Errorf("Email [%v] expected to be [%v] but got [%v]. Err [%v]", test.email, test.expected, valid, err)
		}
	}
}

func TestSubscriptionActive(t *testing.T) {
	ctx, done, _ := aetest.NewContext()
	defer done()

	a := len(Active(ctx))
	if a != 0 {
		t.Fatalf("No active subscriptions should be found but found %v", a)
	}

	s := Subscription{Email: "winston@flores.org"}
	if _, err := s.Subscribe(ctx); err != nil {
		t.Fatalf("Could not subscribe user with email [%v] Err [%v]", s.Email, err)
	}

	a = len(Active(ctx))
	if a != 1 {
		t.Fatalf("Only one subscription should be found but found %v", a)
	}
}

func TestSubscriptionDelete(t *testing.T) {
	ctx, done, _ := aetest.NewContext()
	defer done()
	email := "winston@flores.org"

	s := Subscription{Email: email}

	if _, err := s.Subscribe(ctx); err != nil {
		t.Fatalf("Could not subscribe user with email [%v] Err [%v]", s.Email, err)
	}

	s2 := Subscription{Email: email}
	if err := s2.Delete(ctx); err != nil {
		t.Fatalf("Could not delete subscription ` with email [%v] Err [%v]", s2.Email, err)
	}

	if _, err := s2.Get(ctx); err == nil {
		t.Fatalf("Expected Subscription to be deleted but found one")
	}
}
