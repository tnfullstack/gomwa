package main

import (
	"log"
	"net/http"

	"githum.com/tvn/gomwa/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("about", handlers.About)
	http.HandleFunc("contact", handlers.Contact)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
