package main

import (
	"fmt"
	"gomwa/pkg/handlers"
	"gomwa/pkg/util"
	"log"
	"net/http"
)

func main() {
	port := ":8000"

	util.LocalAssets("assets")

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
