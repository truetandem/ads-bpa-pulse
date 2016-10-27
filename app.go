package main

import (
	"net/http"
	"time"
)

func init() {
	time.Local = time.UTC
	http.HandleFunc("/", Home)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/subscribe", Subscribe)
	http.HandleFunc("/unsubscribe", Unsubscribe)
	http.HandleFunc("/test/mail", TestEmail)
}
