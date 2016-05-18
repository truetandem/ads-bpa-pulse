package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeReturnsOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Home))
	defer server.Close()

	if resp, err := http.DefaultClient.Get(server.URL); err != nil || resp.StatusCode != http.StatusOK {
		t.FailNow()
	}
}
