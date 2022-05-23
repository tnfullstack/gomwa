// Testing Go Web Server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello Gopher, your requested URL is '%s'\n", r.URL.Path)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Number of byte written: %d\n", n)
	})

	http.ListenAndServe(":8080", nil)
}
