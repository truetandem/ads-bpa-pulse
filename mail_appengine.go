// +build appengine

package main

import (
	"log"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"
)

// sendEmail with the AppEngine API.
func sendEmail(r *http.Request, from string, to []string, subject, body, html string) {
	// Use AppEngine to send our thank you cards
	c := appengine.NewContext(r)
	msg := &mail.Message{
		Sender:   from,
		Bcc:      to,
		Subject:  subject,
		Body:     body,
		HTMLBody: html,
	}

	if err := mail.Send(c, msg); err != nil {
		log.Printf("[ERROR] Could not send email: %v\n", err)
	}
}
