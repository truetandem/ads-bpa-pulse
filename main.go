package main

import (
	"log"
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

func main() {
	log.Println("Serving content on", "http://localhost:8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
