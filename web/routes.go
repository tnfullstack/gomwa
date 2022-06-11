package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"githum.com/tvn/gomwa/handlers"
)

func routes() http.Handler {
	mux := mux.NewRouter()

	fs := http.FileServer(http.Dir("../templates/assets"))
	mux.PathPrefix("../templates/assets/").Handler(http.StripPrefix("../templates/assets/", fs))

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/contact", handlers.Contact)

	return mux
}
