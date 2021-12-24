package handlers

import (
	"net/http"

	"githum.com/tnfullstack/gomwa/servwebpage/pkg/render"
)

// Home returns the web app home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About returns the web app about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
