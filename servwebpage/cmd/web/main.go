// Serving web pages
package main

import (
	"fmt"
	"log"
	"net/http"

	"githum.com/tnfullstack/gomwa/servwebpage/pkg/handlers"
)

const (
	webPort = ":8080"
)

// main function is the starting point of the web app.
func main() {

	// HandleFunc listen to web url requests
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting web app on port", webPort)
	// Listen to web server request and serve the pages
	err := http.ListenAndServe(webPort, nil)
	if err != nil {
		log.Println(err)
	}
}
