// +build appengine

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"github.com/PuerkitoBio/goquery"
)

// Update will scrape the source information for solicitations, compare them with
// previously stored results, and notify the subscribers.
func Update(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err := client.Get("https://pages.18f.gov/ads-bpa/")
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Query for the document to scrape
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Loop through the document and parsing solicitations
	solicitations := []Solicitation{}
	doc.Find(".solicitation").Each(func(_ int, dl *goquery.Selection) {
		s := Solicitation{
			Properties: map[string]string{},
		}

		dds := dl.Find("dd")
		dl.Find("dt").Each(func(i int, prop *goquery.Selection) {
			name := strings.TrimSpace(prop.Text())
			value := strings.TrimSpace(dds.Eq(i).Text())

			if name == "Title" {
				s.Title = value
			} else {
				s.Properties[name] = value
			}
		})

		solicitations = append(solicitations, s)
	})

	// Transform for JSON consumtion
	js, err := json.Marshal(solicitations)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return as JSON array
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
