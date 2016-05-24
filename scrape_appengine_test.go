// +build appengine

package main

import (
	"testing"

	"google.golang.org/appengine/aetest"
)

func TestScrape(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	r, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	doc, err := scrape(r, "http://google.com")
	if err != nil {
		t.FailNow()
	}
	if doc == nil {
		t.Fatal("No document")
	}
}

func TestScrapeWithBadAddress(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	r, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	_, err = scrape(r, "gopher://google.com")
	if err == nil {
		t.FailNow()
	}
}
