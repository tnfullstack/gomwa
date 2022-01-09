// Serving web pages
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tvn9/gomwa/web1/pkg/handlers"
)

const (
	webPort = ":8080"
)

// main function is the starting point of the web app.
func main() {
	// Define a http file server to serve static css files
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	// HandleFunc listen to web url requests
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/signup", handlers.Signup)

	fmt.Println("Starting web app on port", webPort)

	// Listen to web server request and serve the pages
	err := http.ListenAndServe(webPort, nil)
	if err != nil {
		log.Println(err)
	}
}
