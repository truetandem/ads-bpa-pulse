package main

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
)

func Subscribe(w http.ResponseWriter, r *http.Request) {

	// Only allow POST requests since we're saving data
	if r.Method != "POST" {
		http.Error(w, "Invalid Method Type", http.StatusMethodNotAllowed)
		return
	}

	// Parse out email param
	r.ParseForm()
	email := r.FormValue("email")

	s := Subscription{Email: email}

	if valid, err := s.ValidEmail(); !valid {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Subscribe using email
	ctx := appengine.NewContext(r)
	if _, err := s.Subscribe(ctx); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Woop
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Message string
		Email   string
	}{
		"Successfully subscribed",
		s.Email,
	})

}
