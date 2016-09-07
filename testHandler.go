package main

import "net/http"

// TestEmail will send a test notification using the mailer.
func TestEmail(w http.ResponseWriter, r *http.Request) {
	err := sendEmail(r, "noreply@ads-bpa-pulse.appspotmail.com", []string{"bryan.allred@truetandem.com"}, "Test", "Test", "Test")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
