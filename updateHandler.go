package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"

	"github.com/PuerkitoBio/goquery"
)

var (
	templateEmailPlain = template.Must(template.ParseFiles("templates/email.txt"))
	templateEmailHTML  = template.Must(template.ParseFiles("templates/email.html"))
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
			o, err := s.Get(ctx)

			// If it never existed give it a time
			if err == datastore.ErrNoSuchEntity {
				s.Modified = time.Now()
			}

			// Check for an existing solicitation with a non-matching
			// checksum of its properties.
			if err == nil && s.Checksum() != o.Checksum() {
				s.Modified = time.Now()
				updates = append(updates, s)
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

	// Notify subscribers of any updates
	if len(updates) > 0 {
		subscriptions := Active(ctx)
		if len(subscriptions) > 0 {
			var plain bytes.Buffer
			err = templateEmailPlain.Execute(&plain, updates)

			var html bytes.Buffer
			err = templateEmailHTML.Execute(&html, updates)

			if err == nil {
				sendEmail(
					r,
					"no-reply@truetandem.com",
					subscriptions,
					"A pulse was identified for ADS-BPA",
					plain.String(),
					html.String())
			}
		}
	}

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
