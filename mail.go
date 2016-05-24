// +build !appengine

package main

import (
	"net/http"
	"net/smtp"
)

// sendEmail without the AppEngine libraries.
func sendEmail(r *http.Request, from string, to []string, subject, body, html string) error {
	auth := smtp.PlainAuth("", "", "", "test.smtp.org")
	return smtp.SendMail("test.smtp.org:25", auth, from, to, []byte(body))
}
