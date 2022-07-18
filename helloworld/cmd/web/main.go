package main

import (
	"fmt"
	"gomwa/helloworld/pkg/handlers"
	"log"
	"net/http"
)

const portNum = ":8080"

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting server on port number:", portNum)
	log.Fatal(http.ListenAndServe(portNum, nil))
}
