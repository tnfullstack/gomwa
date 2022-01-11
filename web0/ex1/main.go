package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	// ListionAndServe
	http.ListenAndServe(":8080", nil)

}

// handler
func handler(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello, Gopher Again!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number of bytes written %d\n", n)
}
