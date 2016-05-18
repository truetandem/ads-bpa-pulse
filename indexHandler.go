package main

import (
	"html/template"
	"net/http"
)

var (
	templates = template.Must(template.ParseFiles("index.html"))
)

// Home page where user can subscribe and unsubscribe to BPA updates.
func Home(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
