// Testing Go Web Server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, World!, your request URL is '%s'\n", r.URL.Path)
	if err != nil {
		log.Fatal(err)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is about page!")
}

func Math(w http.ResponseWriter, r *http.Request) {
	sum := addNums(5, 8)
	_, err := fmt.Fprintf(w, "The sum of 5 + 8 = %d\n", sum)
	if err != nil {
		log.Fatal(err)
	}
}

func addNums(x, y int) int {
	return x + y
}

func main() {

	port := ":8080"

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/add", Math)

	fmt.Println("Starting Server at port", port)
	http.ListenAndServe(port, nil)
}
