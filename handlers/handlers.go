package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpls *template.Template

func init() {
	tmpls = template.Must(template.ParseGlob("../templates/*.html"))
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := tmpls.ExecuteTemplate(w, tmpl, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
