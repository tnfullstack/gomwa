// Serving web pages
package main

import (
	"fmt"
	"log"
	"net/http"
	"web1/pkg/handlers"

	"github.com/gorilla/mux"
)

const (
	webPort = ":9999"
)

// main function is the starting point of the web app.
func main() {
	// Define a gorilla mux router
	r := mux.NewRouter()
	
	// Define a http file server to serve static css files (this block use without mux)
	// fs := http.FileServer(http.Dir("../assets"))
	// http.Handle("/assets/", http.StripPrefix("/assets", fs))
	// This will serve files under http://localhost:8000/static/<filename> (use with mux.NewServer)

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("../assets"))))

	// HandleFunc listen to web url requests
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/login", handlers.Login)
	r.HandleFunc("/contacts", handlers.Contacts)
	r.HandleFunc("/about", handlers.About)
	r.HandleFunc("/signup", handlers.Signup)
	http.Handle("/", r)
	fmt.Println("Starting web app on port", webPort)

	// Listen to web server request and serve the pages
	err := http.ListenAndServe(webPort, r)
	if err != nil {
		log.Println(err)
	}
}
