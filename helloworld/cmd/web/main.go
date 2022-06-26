package main

import (
	"github.com/tvn9/gomwa/helloworld/pkg/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(":8080", nil)
}
