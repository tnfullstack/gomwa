package handler

import (
	"net/http"

	"github.com/tvn9/gomwa/web3/render"
)

// Home
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.page.html")
}

// About
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "about.page.html")
}
