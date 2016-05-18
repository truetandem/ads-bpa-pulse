// +build !appengine

package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func scrape(r *http.Request, uri string) (*goquery.Document, error) {
	return goquery.NewDocument(uri)
}
