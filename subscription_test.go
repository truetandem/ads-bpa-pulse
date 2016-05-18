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
		t.Fatal("Unable to Save Subscription [%v]\n", err)
	}

	if key == nil {
		t.Fatal("Key was not set")
	}

	if key.StringID() != email {
		t.Fatal("Key string id does not match email")
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
	}

	for _, test := range emailTests {
		s := Subscription{Email: test.email}
		if valid, err := s.ValidEmail(); valid != test.expected {
			t.Errorf("Email [%v] expected to be [%v] but got [%v]. Err [%v]", test.email, test.expected, valid, err)
		}
	}
}
