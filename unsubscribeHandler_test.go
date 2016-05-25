package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

func TestUnsubscribe(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	tests := []struct {
		Method       string
		ExpectedCode int
		UrlValues    url.Values
	}{
		{"POST", http.StatusOK, url.Values{"email": {"johnsflores@gmail.com"}}},
	}

	for _, test := range tests {
		if r, err := inst.NewRequest(test.Method, "/unsubscribe", nil); err == nil {
			// Create sample subscription that we'll delete later
			s := Subscription{Email: "johnsflores@gmail.com"}
			ctx := appengine.NewContext(r)
			if _, err := s.Subscribe(ctx); err != nil {
				t.Fatalf("Unable to create Subscription.")
			}
			r.Form = test.UrlValues
			w := httptest.NewRecorder()
			Unsubscribe(w, r)
			if w.Code != test.ExpectedCode {
				t.Errorf("Expected Status Code [%v] Got [%v]", test.ExpectedCode, w.Code)
			}

			w = httptest.NewRecorder()
			Unsubscribe(w, r)
			if w.Code != 500 {
				t.Errorf("Expected Status Code [%v] Got [%v]", 500, w.Code)
			}
		} else {

			t.Errorf("Unable to create request using Method [%v] Expected Code [%v] UrlValues [%v]", test.Method, test.ExpectedCode, test.UrlValues)
		}
	}
}

func TestInvalidUnsubscribe(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	tests := []struct {
		Method       string
		ExpectedCode int
		URLValues    url.Values
	}{
		{"GET", http.StatusMethodNotAllowed, nil},
		{"POST", http.StatusInternalServerError, nil},
	}

	for _, test := range tests {
		if r, err := inst.NewRequest(test.Method, "/unsubscribe", nil); err == nil {
			r.Form = test.URLValues
			w := httptest.NewRecorder()
			Unsubscribe(w, r)
			if w.Code != test.ExpectedCode {
				t.Errorf("Expected Status Code [%v] Got [%v]", w.Code, test.ExpectedCode)
			}
		} else {
			t.Errorf("Unable to create request using Method [%v] Expected Code [%v] UrlValues [%v]", test.Method, test.ExpectedCode, test.URLValues)
		}
	}
}
