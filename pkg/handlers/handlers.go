package handlers

import (
	"gomwa/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "index.tmpl.html")

}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.tmpl.html")
}
