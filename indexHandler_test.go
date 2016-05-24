package main

import (
	"html/template"
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

func TestHomeBadTemplate(t *testing.T) {
	templateIndex = template.New("bad")
	server := httptest.NewServer(http.HandlerFunc(Home))
	defer server.Close()

	resp, _ := http.DefaultClient.Get(server.URL)
	if resp.StatusCode != http.StatusInternalServerError {
		t.FailNow()
	}
}
