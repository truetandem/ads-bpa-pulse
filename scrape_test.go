// +build !appengine

package main

import "testing"

func TestScrape(t *testing.T) {
	doc, err := scrape(nil, "http://google.com")
	if err != nil {
		t.FailNow()
	}
	if doc == nil {
		t.Fatal("No document")
	}
}

func TestScrapeWithBadAddress(t *testing.T) {
	_, err := scrape(nil, "gopher://google.com")
	if err == nil {
		t.FailNow()
	}
}
