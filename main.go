// +build !appengine

package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Serving content on", "http://localhost:8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
