package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"google.golang.org/appengine"

	"github.com/PuerkitoBio/goquery"
)

// Update will scrape the source information for solicitations, compare them with
// previously stored results, and notify the subscribers.
func Update(w http.ResponseWriter, r *http.Request) {
	// Query for the document to scrape
	doc, err := scrape(r, "https://pages.18f.gov/ads-bpa/")
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Loop through the document and parsing solicitations
	updates := []Solicitation{}
	solicitations := []Solicitation{}
	ctx := appengine.NewContext(r)
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

		// If there is no title then we disregard the item
		if s.Title != "" {
			// Check for an existing solicitation with a non-matching
			// checksum of its properties.
			u, err := s.Get(ctx)
			if err == nil && s.Checksum() != u.Checksum() {
				updates = append(updates, u)
			}

			// Save the solicitation
			err = s.Save(ctx)
			if err != nil {
				log.Fatal(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			solicitations = append(solicitations, s)
		}
	})

	// Transform for JSON consumption
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
