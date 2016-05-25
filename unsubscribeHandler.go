package main

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
)

// Unsubscribe handles a request to be removed from the notification listing.
func Unsubscribe(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests since we're saving data
	if r.Method != "POST" {
		http.Error(w, "Invalid Method Type", http.StatusMethodNotAllowed)
		return
	}

	// Parse out email param
	_ = r.ParseForm()
	email := r.FormValue("email")

	s := Subscription{Email: email}
	if valid, err := s.ValidEmail(); !valid {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Unsubscribe using email
	ctx := appengine.NewContext(r)
	if err := s.Unsubscribe(ctx); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Woop
	w.Header().Set("Content-Type", "application/json")
	m := message{
		"Successfully unsubscribed",
		s.Email,
	}

	err := json.NewEncoder(w).Encode(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
