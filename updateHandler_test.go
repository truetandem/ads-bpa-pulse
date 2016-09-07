// +build !appengine

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestUpdateReturnsOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Update))
	defer server.Close()

	if resp, err := http.DefaultClient.Get(server.URL); err != nil || resp.StatusCode != http.StatusOK {
		t.FailNow()
	}
}

func TestUpdateReturnsSolicitations(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Update))
	defer server.Close()

	resp, err := http.DefaultClient.Get(server.URL)
	if err != nil || resp.StatusCode != http.StatusOK {
		t.FailNow()
	}
	defer resp.Body.Close()

	dump, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.FailNow()
	}

	var solicitations []Solicitation
	err = json.Unmarshal(dump, &solicitations)
	if err != nil {
		t.FailNow()
	}
}

func TestUpdateFindsUpdatedPropertyValues(t *testing.T) {
	tests := []struct {
		file          string
		updates       int
		solicitations int
	}{
		{"sample-feed.html", 5, 5},
		{"sample-feed-award.html", 1, 5},
	}

	for _, test := range tests {
		file, err := os.Open(test.file)
		if err != nil {
			t.Errorf("Could not open '%s'", test.file)
		}

		doc, err := goquery.NewDocumentFromReader(file)
		if err != nil {
			t.Errorf("Could not parse '%s'", test.file)
		}

		updates, solicitations, err := parseDocument(nil, doc)
		if err != nil && len(updates) != test.updates && len(solicitations) != test.solicitations {
			t.FailNow()
		}
	}
}

func TestUpdateFindsNewSolicitation(t *testing.T) {
	tests := []struct {
		file          string
		updates       int
		solicitations int
	}{
		{"testdata/sample-feed.html", 5, 5},
		{"testdata/sample-feed-new.html", 1, 6},
	}

	for _, test := range tests {
		file, err := os.Open(test.file)
		if err != nil {
			t.Errorf("Could not open '%s'", test.file)
		}

		doc, err := goquery.NewDocumentFromReader(file)
		if err != nil {
			t.Errorf("Could not parse '%s'", test.file)
		}

		updates, solicitations, err := parseDocument(nil, doc)
		if err != nil {
			t.Errorf("Error during parsing: %v", err)
		} else if len(solicitations) != test.solicitations {
			t.Errorf("In %s expected %d but received %d solicitations", test.file, test.solicitations, len(solicitations))
		} else if len(updates) != test.updates {
			t.Errorf("In %s expected %d but received %d updates", test.file, test.updates, len(updates))
		}
	}
}
