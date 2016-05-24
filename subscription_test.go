package main

import (
	"fmt"
	"testing"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

func TestSubscriptionSave(t *testing.T) {
	ctx, done, _ := aetest.NewContext()
	defer done()

	var tests = []struct {
		Email         string
		ExpectedError error
	}{
		{Email: "johnsflores@gmail.com", ExpectedError: nil},
		{Email: "john", ExpectedError: ErrEmailInvalid},
		{Email: "", ExpectedError: ErrEmailRequired},
	}

	for _, test := range tests {

		s := Subscription{Email: test.Email}
		_, err := s.Save(ctx)

		if err != test.ExpectedError {
			t.Fatalf("Expected Error [%v] but got [%v]", test.ExpectedError, err)
		}
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

	s3 := Subscription{Email: "nonexistant@mail.com"}
	if err := s3.Unsubscribe(ctx); err != ErrSubscriptionDoesNotExist {
		t.FailNow()
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
	options := aetest.Options{
		AppID: "testapp",
		StronglyConsistentDatastore: true,
	}
	inst, err := aetest.NewInstance(&options)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	r, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	ctx := appengine.NewContext(r)
	a := len(Active(ctx))
	if a != 0 {
		t.Fatalf("No active subscriptions should be found but found %v", a)
	}

	s := Subscription{Email: "winston@flores.org"}
	if _, err := s.Save(ctx); err != nil {
		t.Fatalf("Could not subscribe user with email [%v] Err [%v]", s.Email, err)
	}

	ctx = appengine.NewContext(r)
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

func TestSubscriptionString(t *testing.T) {
	s := Subscription{
		Email: "string@test.com",
	}
	expected := fmt.Sprintf("Subscription - Email:  [%v] Modfied: [%v]", s.Email, s.Modified)

	if s.String() != expected {
		t.FailNow()
	}
}
