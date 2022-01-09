package handlers

import (
	"net/http"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/*.tmpl"))
}

// Home returns the web app home page
func Home(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.tmpl", nil)
}

// About returns the web app about page
func About(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.tmpl", nil)
}

// About returns the web app about page
func Login(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "login.tmpl", nil)
}

// About returns the web app about page
func Signup(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "signup.tmpl", nil)
}
