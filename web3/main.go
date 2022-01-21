package main

import (
	"net/http"

	"github.com/tvn9/gomwa/web3/handler"
)

func main() {

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	srv := http.Server{
		Addr: ":8080",
	}

	srv.ListenAndServe()
}
