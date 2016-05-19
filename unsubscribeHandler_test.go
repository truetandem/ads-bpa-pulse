package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

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
		{"GET", http.StatusMethodNotAllowed, nil},
		{"POST", http.StatusInternalServerError, nil},
		{"POST", http.StatusOK, url.Values{"email": {"johnsflores@gmail.com"}}},
	}

	// Create sample subscription that we'll delete later
	s := Subscription{Email: "johnsflores@gmail.com"}
	ctx, done, _ := aetest.NewContext()
	defer done()
	if _, err := s.Subscribe(ctx); err != nil {
		t.Fatalf("Unable to create Subscription.")
	}

	for _, test := range tests {
		if r, err := inst.NewRequest(test.Method, "/unsubscribe", nil); err == nil {
			r.Form = test.UrlValues
			w := httptest.NewRecorder()
			Subscribe(w, r)
			if w.Code != test.ExpectedCode {
				t.Errorf("Expected Status Code [%v] Got [%v]", w.Code, test.ExpectedCode)
			}
		} else {

			t.Errorf("Unable to create request using Method [%v] Expected Code [%v] UrlValues [%v]", test.Method, test.ExpectedCode, test.UrlValues)
		}
	}
}
