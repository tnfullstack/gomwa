package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	srv := http.Server{
		Addr: ":8080",
	}

	srv.ListenAndServe()
}
