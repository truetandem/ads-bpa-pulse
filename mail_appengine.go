// +build appengine

package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"
)

// sendEmail with the AppEngine API.
func sendEmail(r *http.Request, from string, to []string, subject, body, html string) error {
	// Use AppEngine to send our thank you cards
	c := appengine.NewContext(r)
	msg := &mail.Message{
		Sender:   from,
		Bcc:      to,
		Subject:  subject,
		Body:     body,
		HTMLBody: html,
	}

	return mail.Send(c, msg)
}
