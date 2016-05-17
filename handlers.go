package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Update will scrape the source information for solicitations, compare them with
// previously stored results, and notify the subscribers.
func Update(w http.ResponseWriter, r *http.Request) {
	// Query for the document to scrape
	doc, err := goquery.NewDocument("https://pages.18f.gov/ads-bpa/")
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
	w.Write(js)
}
