package main

import (
	"fmt"
	"net/http"

	"github.com/tvn9/gomwa/pkg/handlers"
)

const portNum = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/contact", handlers.Contact)

	fmt.Println("Starting server on port:", portNum)
	_ = http.ListenAndServe(portNum, nil)
}
