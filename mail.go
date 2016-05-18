// +build !appengine

package main

import (
	"log"
	"net/http"
	"net/smtp"
)

// sendEmail without the AppEngine libraries.
func sendEmail(r *http.Request, from string, to []string, subject, body, html string) {
	auth := smtp.PlainAuth("", "", "", "test.smtp.org")
	if err := smtp.SendMail("test.smtp.org:25", auth, from, to, []byte(body)); err != nil {
		log.Println(err)
	}
}
