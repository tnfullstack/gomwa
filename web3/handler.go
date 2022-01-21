package main

import (
	"net/http"
)

// Home
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home.page.html")
}

// About
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about.page.html")
}
