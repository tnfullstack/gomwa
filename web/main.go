package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNum = ":8080"

func main() {
	mux := routes()

	fmt.Printf("Starting server on port %s\n", portNum)
	log.Fatal(http.ListenAndServe(portNum, mux))
}
