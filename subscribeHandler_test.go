package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"google.golang.org/appengine/aetest"
)

func TestSubscribe(t *testing.T) {
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
		{"POST", http.StatusInternalServerError, url.Values{"email": {"johnsflores@gmail.com"}}},
	}

	for _, test := range tests {
		if r, err := inst.NewRequest(test.Method, "/subscribe", nil); err == nil {
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
