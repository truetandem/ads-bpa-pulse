// +build appengine

package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"github.com/PuerkitoBio/goquery"
)

func scrape(r *http.Request, uri string) (*goquery.Document, error) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)

	resp, err := client.Get(uri)
	if err != nil {
		return nil, err
	}

	return goquery.NewDocumentFromResponse(resp)
}
