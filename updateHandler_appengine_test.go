// +build appengine

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

func TestUpdateReturnsOK(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	r, err := inst.NewRequest("GET", "/update", nil)
	if err != nil {
		t.FailNow()
	}

	w := httptest.NewRecorder()
	Update(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Expected Status Code [%v] Got [%v]", w.Code, http.StatusOK)
	}
}

func TestUpdateReturnsSolicitations(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	r, err := inst.NewRequest("GET", "/update", nil)
	if err != nil {
		t.FailNow()
	}

	w := httptest.NewRecorder()
	Update(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Expected Status Code [%v] Got [%v]", w.Code, http.StatusOK)
	}

	dump, err := ioutil.ReadAll(w.Body)
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
		{"testdata/sample-feed.html", 5, 5},
		{"testdata/sample-feed-award.html", 1, 5},
	}

	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	if request, err := inst.NewRequest("POST", "/", nil); err == nil {
		ctx := appengine.NewContext(request)
		for _, test := range tests {
			file, err := os.Open(test.file)
			if err != nil {
				t.Errorf("Could not open '%s'", test.file)
			}

			doc, err := goquery.NewDocumentFromReader(file)
			if err != nil {
				t.Errorf("Could not parse '%s'", test.file)
			}

			updates, solicitations, err := parseDocument(ctx, doc)
			if err != nil {
				t.Errorf("Error during parsing: %v", err)
			} else if len(solicitations) != test.solicitations {
				t.Errorf("In %s expected %d but received %d solicitations", test.file, test.solicitations, len(solicitations))
			} else if len(updates) != test.updates {
				t.Errorf("In %s expected %d but received %d updates", test.file, test.updates, len(updates))
			}
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

	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	if request, err := inst.NewRequest("POST", "/", nil); err == nil {
		ctx := appengine.NewContext(request)
		for _, test := range tests {
			file, err := os.Open(test.file)
			if err != nil {
				t.Errorf("Could not open '%s'", test.file)
			}

			doc, err := goquery.NewDocumentFromReader(file)
			if err != nil {
				t.Errorf("Could not parse '%s'", test.file)
			}

			updates, solicitations, err := parseDocument(ctx, doc)
			if err != nil {
				t.Errorf("Error during parsing: %v", err)
			} else if len(solicitations) != test.solicitations {
				t.Errorf("In %s expected %d but received %d solicitations", test.file, test.solicitations, len(solicitations))
			} else if len(updates) != test.updates {
				t.Errorf("In %s expected %d but received %d updates", test.file, test.updates, len(updates))
			}
		}
	}
}
